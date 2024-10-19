package ih2024

import (
	"context"
	"encoding/json"

	"github.com/Rhymond/go-money"
	"github.com/ThreeDotsLabs/watermill"
	"github.com/ThreeDotsLabs/watermill/message"
	"github.com/ThreeDotsLabs/watermill/pubsub/gochannel"
	"github.com/pkg/errors"
	"github.com/shopd/shopd/go/model"
	"github.com/shopd/shopd/go/msg"
	"github.com/shopd/shopd/go/share"
)

// ProcessMsg handler
func (ph *Handler) ProcessMsg(m *message.Message) error {
	var wh Webhook
	err := json.Unmarshal(m.Payload, &wh)
	if err != nil {
		return errors.WithStack(err)
	}

	qx, err := ph.s.DomainQX(context.Background())
	if err != nil {
		return err
	}

	switch wh.Event.Result.String {
	case GrantAccepted:
		// Parse payload
		orders, err := model.NewOrders(qx.QX, model.OrdersByOrderNo(wh.OrderNo))
		if err != nil {
			return err
		}
		if len(orders) != 1 {
			return ErrOrderNo(wh.OrderNo)
		}
		order := orders[0]

		// Save the transaction and link order
		err = model.SaveOrderTransaction(qx.QX, model.SaveOrderTransactionParams{
			OrderID: order.OrderID.String,
			UserID:  order.UserID.String,
			Amount:  wh.Event.Amount.Int64,
			// TODO Make Descr configurable?
			Descr: "Interledger Hackathon 2024",
		})
		if err != nil {
			return err
		}

		// Paid in full?
		// Use go-money everywhere money calculations are done,
		// even though the int64 values could be compared directly,
		// makes it easier to search for code like this
		tranAmount := money.New(
			wh.Event.Amount.Int64*100, qx.DomainConfig.Currency())
		orderTotal := money.New(
			order.Totals.Total.Int64, qx.DomainConfig.Currency())
		paid, err := tranAmount.GreaterThanOrEqual(orderTotal)
		if err != nil {
			return err
		}

		// Confirm order
		order.Confirm(qx.QX, model.OrderConfirmParams{
			ModID: model.SystemUserID,
			Paid:  paid,
		})

	case GrantRejected:
		return ErrNotImplemented

	}
	return nil
}

// registerPubSub registers publishers and subscribers
func (ph *Handler) registerPubSub() error {
	// Publisher for use with the webhooks route
	fields := watermill.LogFields{}
	fields.Add(watermill.LogFields{
		share.ParamPlugin: ph.Name(),
	})
	pubSub := gochannel.NewGoChannel(
		gochannel.Config{},
		msg.NewLogger(fields),
	)
	err := ph.s.Msg().RegisterPublisher(ph.Name(), pubSub)
	if err != nil {
		return err
	}

	// Subscriber for processing webhooks
	_, err = ph.s.Msg().RegisterSubscriber(
		msg.NewTopic(ph.Name()), ph.ProcessMsg)
	if err != nil {
		return err
	}

	return nil
}
