package ih2024_test

import (
	"context"
	"encoding/json"
	"fmt"
	"sync"
	"testing"

	"github.com/Rhymond/go-money"
	"github.com/ThreeDotsLabs/watermill"
	"github.com/ThreeDotsLabs/watermill/message"
	"github.com/ThreeDotsLabs/watermill/pubsub/gochannel"
	"github.com/mozey/ft"
	"github.com/rs/zerolog/log"
	"github.com/shopd/shopd/go/dbgen"
	"github.com/shopd/shopd/go/model"
	"github.com/shopd/shopd/go/msg"
	"github.com/shopd/shopd/go/plugin/ih2024"
	"github.com/shopd/shopd/go/plugin/ih2024/config"
	"github.com/shopd/shopd/go/share"
	"github.com/shopd/shopd/go/testutil"
)

func TestProcessMsg(t *testing.T) {
	is, s, ps := testutil.SetupPluginServices(t)

	pConf := config.New()
	ph := ih2024.New(pConf, ps)

	// ...........................................................................
	// pubsub

	// wg is used to wait for async messages to process
	var wg sync.WaitGroup

	topic := msg.NewTopic(t.Name())

	pubSub := gochannel.NewGoChannel(
		gochannel.Config{},
		msg.NewLogger(watermill.LogFields{}),
	)

	// Register publisher
	err := s.Msg().RegisterPublisher(topic, pubSub)
	is.NoErr(err)

	// Register subscriber
	subscriberID, err := s.Msg().RegisterSubscriber(topic,
		func(m *message.Message) error {
			log.Info().Msg("Subscriber")
			err := ph.ProcessMsg(m)
			if err != nil {
				log.Error().Stack().Err(err).Msg("")
			}
			wg.Done()
			return nil
		})
	is.NoErr(err)

	// Start subscriber
	err = s.Msg().StartSubscriber(subscriberID)
	is.NoErr(err)

	qx, err := s.DomainQX(context.Background())
	is.NoErr(err)

	// ...........................................................................
	// Clean
	_, err = qx.NamedExec("delete from orders where order_id like :order_id",
		map[string]any{
			"order_id": fmt.Sprintf("%s%%", t.Name()),
		})
	is.NoErr(err)
	_, err = qx.NamedExec("delete from order_line where order_id like :order_id",
		map[string]any{
			"order_id": fmt.Sprintf("%s%%", t.Name()),
		})
	is.NoErr(err)
	_, err = qx.NamedExec("delete from cat where sku like :sku",
		map[string]any{
			"sku": fmt.Sprintf("%s%%", t.Name()),
		})
	is.NoErr(err)

	// ...........................................................................
	// Create test data

	// Inventory
	sku := t.Name()
	_, err = qx.CatUpsert(qx.Context(), dbgen.CatUpsertParams{
		ColMap: dbgen.CatUpsertCols(),
		Rows: []dbgen.Cat{{
			SKU:   sku,
			State: share.CSStock,
			ModID: model.SystemUserID,
		}},
	})
	is.NoErr(err)

	_, err = qx.CatQtyAdjustment(qx.Context(), dbgen.CatQtyAdjustmentParams{
		Adjustments: []share.CatalogQtyAdjustment{{
			Sku:      sku,
			Depot:    "",
			QtyDelta: 1,
		}},
	})
	is.NoErr(err)

	// Order with state processing
	cart := model.NewCartFromItems([]model.CartFromItemsParams{{
		Sku: sku,
		Item: share.CartItem{
			Qty: ft.NIntFrom(1),
		},
	}})

	user, _, err := testutil.CreateAccessToken(t, s, share.UserRoleCustomer)
	is.NoErr(err)

	qx.DomainConfig.SetCountry(share.CountryZAF)
	qx.DomainConfig.SetCurrency(money.ZAR)
	qx.DomainConfig.SetTaxCalculation(share.TaxByBasket)
	order, err := cart.Quote(qx, model.QuoteParams{
		UserID: user.UserID,
	})
	is.NoErr(err)

	err = order.Reserve(qx.QX, model.OrderReserveParams{
		ModID: model.SystemUserID,
	})
	is.NoErr(err)
	log.Info().Interface("order", order).Msg("")

	// ...........................................................................
	// Simulate payment workflow

	wg.Add(1) // Publish one message

	// Payload must cause the message handler to confirm the order
	wh := ih2024.Webhook{}
	wh.Event.Result = ft.NStringFrom(ih2024.GrantAccepted)
	wh.OrderNo = order.OrderNo.String
	wh.Event.Amount = ft.NIntFrom(101)
	log.Info().Interface("webhook", wh).Msg("")
	payload, err := json.Marshal(wh)
	is.NoErr(err)
	// Publish message directly, don't use webhook route.
	s.Msg().Publish(topic, msg.NewMessage(payload))

	wg.Wait() // Wait for messages to process

	// Verify order
	var orderRows []dbgen.Orders
	err = qx.NamedSelect(&orderRows,
		"select * from orders where order_id = :order_id", map[string]any{
			"order_id": order.OrderID.String,
		})
	is.NoErr(err)
	is.Equal(len(orderRows), 1) // Expected an order
	is.Equal(orderRows[0].OrderID, order.OrderID.String)
	is.Equal(orderRows[0].State, string(share.OSConfirmed))
	is.Equal(orderRows[0].Paid, int64(1))

	trans, err := model.NewTransaction(
		qx.QX, model.TransactionsByOrderID(order.OrderID.String))
	is.NoErr(err)
	is.Equal(len(trans), 1)                                // Expected a transaction
	is.Equal(trans[0].Amount.Int64, wh.Event.Amount.Int64) // Amount must match webhook
}
