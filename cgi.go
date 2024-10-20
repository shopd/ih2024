package ih2024

import (
	"context"
	"encoding/json"
	"fmt"
	"os/exec"
	"path/filepath"

	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"
	"github.com/shopd/shopd/go/fileutil"
	"github.com/shopd/shopd/go/model"
)

type Redirect struct {
	RedirectURL string
	ContinueURI string
	AccessToken string
	QuoteID     string
}

type Continue struct {
	Message string
}

type NewRedirectParams struct {
	SuccessURL string
	Nonce      string
	Amount     int64
	OrderID    string
	OrderNo    string
}

// paymentRedirect meta data file name.
// In prod the payment meta data must be stored in the order_config table
const paymentRedirect = "redirect.json"

// paymentParams meta data file name.
// In prod this would also be stored on a config table
const paymentParams = "params.json"

// paymentContinue meta data file name.
// In prod this would also be stored on a config table
const paymentContinue = "continue.json"

func (ph *Handler) NewRedirect(params NewRedirectParams) (redirect *Redirect, err error) {
	inWalletAddressURL := ph.conf.Ih2024InWalletAddressUrl()
	// TODO outWalletAddressURL should be provided by the customer after selecting payment method
	outWalletAddressURL := ph.conf.Ih2024OutWalletAddressUrl()
	keyID := ph.conf.Ih2024KeyId()
	privateKey := ph.conf.Ih2024PrivateKey()

	cmd := exec.Command("node", "redirect.js")
	cmd.Env = append(
		cmd.Env, fmt.Sprintf("APP_IH2024_IN_WALLET_ADDRESS_URL=%s", inWalletAddressURL))
	cmd.Env = append(
		cmd.Env, fmt.Sprintf("APP_IH2024_OUT_WALLET_ADDRESS_URL=%s", outWalletAddressURL))
	cmd.Env = append(
		cmd.Env, fmt.Sprintf("APP_IH2024_KEY_ID=%s", keyID))
	cmd.Env = append(
		cmd.Env, fmt.Sprintf("APP_IH2024_PRIVATE_KEY=%s", privateKey))
	cmd.Env = append(
		cmd.Env, fmt.Sprintf("APP_IH2024_SUCCESS_URL=%s", params.SuccessURL))
	cmd.Env = append(
		cmd.Env, fmt.Sprintf("APP_IH2024_NONCE=%s", params.Nonce))
	cmd.Env = append(
		cmd.Env, fmt.Sprintf("APP_IH2024_AMOUNT=%d", params.Amount))
	dir := filepath.Join(ph.conf.Dir(), "go", "plugin", "ih2024")
	cmd.Dir = dir

	out, err := cmd.CombinedOutput()
	if err != nil {
		log.Error().Str("out", string(out)).Msg("")
		return redirect, errors.WithStack(err)
	}

	err = json.Unmarshal(out, &redirect)
	if err != nil {
		log.Error().Str("out", string(out)).Msg("")
		return redirect, errors.WithStack(err)
	}

	// Don't unmarshal before writing,
	// something weird happening with escape chars
	err = fileutil.WriteBytes(filepath.Join(dir, paymentRedirect), out)
	if err != nil {
		return redirect, errors.WithStack(err)
	}

	// Write params to file
	b, err := json.Marshal(params)
	if err != nil {
		return redirect, errors.WithStack(err)
	}
	err = fileutil.WriteBytes(filepath.Join(dir, paymentParams), b)
	if err != nil {
		return redirect, errors.WithStack(err)
	}

	return redirect, nil
}

// ContinueGrant func for hackathon demo...
func (ph *Handler) ContinueGrant() (redirect *Redirect, err error) {
	// ...........................................................................
	// Read payment meta data written to file by NewRedirect
	dir := filepath.Join(ph.conf.Dir(), "go", "plugin", "ih2024")
	b, err := fileutil.ReadAll(filepath.Join(dir, paymentRedirect))
	if err != nil {
		return redirect, err
	}
	err = json.Unmarshal(b, &redirect)
	if err != nil {
		return redirect, errors.WithStack(err)
	}
	log.Debug().Interface("redirect", redirect).Msg("")

	b, err = fileutil.ReadAll(paymentParams)
	if err != nil {
		return redirect, err
	}
	var params NewRedirectParams
	err = json.Unmarshal(b, &params)
	if err != nil {
		return redirect, errors.WithStack(err)
	}
	log.Debug().Interface("params", params).Msg("")

	// ...........................................................................
	// Continue grant

	inWalletAddressURL := ph.conf.Ih2024InWalletAddressUrl()
	// TODO outWalletAddressURL should be provided by the customer after selecting payment method
	outWalletAddressURL := ph.conf.Ih2024OutWalletAddressUrl()
	keyID := ph.conf.Ih2024KeyId()
	privateKey := ph.conf.Ih2024PrivateKey()

	cmd := exec.Command("node", "continue.js")
	cmd.Env = append(
		cmd.Env, fmt.Sprintf("APP_IH2024_IN_WALLET_ADDRESS_URL=%s", inWalletAddressURL))
	cmd.Env = append(
		cmd.Env, fmt.Sprintf("APP_IH2024_OUT_WALLET_ADDRESS_URL=%s", outWalletAddressURL))
	cmd.Env = append(
		cmd.Env, fmt.Sprintf("APP_IH2024_KEY_ID=%s", keyID))
	cmd.Env = append(
		cmd.Env, fmt.Sprintf("APP_IH2024_PRIVATE_KEY=%s", privateKey))
	cmd.Env = append(
		cmd.Env, fmt.Sprintf("APP_IH2024_CONTINUE_URI=%s", redirect.ContinueURI))
	cmd.Env = append(
		cmd.Env, fmt.Sprintf("APP_IH2024_CONTINUE_ACCESS_TOKEN=%s", redirect.AccessToken))
	cmd.Env = append(
		cmd.Env, fmt.Sprintf("APP_IH2024_QUOTE_ID=%s", redirect.QuoteID))
	cmd.Dir = dir

	out, err := cmd.CombinedOutput()
	if err != nil {
		log.Error().Str("out", string(out)).Msg("")
		return redirect, errors.WithStack(err)
	}

	var cont Continue
	err = json.Unmarshal(out, &cont)
	if err != nil {
		log.Error().Str("out", string(out)).Msg("")
		return redirect, errors.WithStack(err)
	}
	err = fileutil.WriteBytes(filepath.Join(dir, paymentContinue), out)
	if err != nil {
		return redirect, errors.WithStack(err)
	}

	// ...........................................................................
	// Confirm order

	qx, err := ph.s.DomainQX(context.Background())
	if err != nil {
		return redirect, err
	}

	orders, err := model.NewOrders(qx.QX, model.OrdersByOrderNo(params.OrderNo))
	if err != nil {
		return redirect, err
	}
	if len(orders) != 1 {
		return redirect, ErrOrderNo(params.OrderNo)
	}
	order := orders[0]

	// Save the transaction and link order
	err = model.SaveOrderTransaction(qx.QX, model.SaveOrderTransactionParams{
		OrderID: order.OrderID.String,
		UserID:  order.UserID.String,
		Amount:  params.Amount,
		Descr:   "Interledger Hackathon 2024",
	})
	if err != nil {
		return redirect, err
	}

	// TODO Skipping this check for Hackathon demo
	// Paid in full?
	// Use go-money everywhere money calculations are done,
	// even though the int64 values could be compared directly,
	// makes it easier to search for code like this
	// tranAmount := money.New(
	// 	redirect.Amount, qx.DomainConfig.Currency())
	// orderTotal := money.New(
	// 	order.Totals.Total.Int64, qx.DomainConfig.Currency())
	// paid, err := tranAmount.GreaterThanOrEqual(orderTotal)
	// if err != nil {
	// 	return redirect, err
	// }

	// Confirm order
	order.Confirm(qx.QX, model.OrderConfirmParams{
		ModID: model.SystemUserID,
		Paid:  true,
	})

	return redirect, nil
}
