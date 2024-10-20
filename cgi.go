package ih2024

import (
	"encoding/json"
	"fmt"
	"os/exec"

	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"
	"github.com/shopd/shopd/go/fileutil"
)

type Redirect struct {
	RedirectURL string
	ContinueURI string
	AccessToken string
	QuoteID     string
	Amount      int64
	OrderID     string
}

type NewRedirectParams struct {
	SuccessURL string
	Nonce      string
	Amount     int64
	OrderID    string
}

// paymentMetaData file name.
// In prod the payment meta data must be stored in the order_config table
const paymentMetaData = "redirect.json"

// paymentResult file name.
// In prod this would also be stored on a config table
const paymentResult = "continue.json"

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
	err = fileutil.WriteBytes(paymentMetaData, out)
	if err != nil {
		return redirect, errors.WithStack(err)
	}

	return redirect, nil
}

// ContinueGrant func for hackathon demo...
func (ph *Handler) ContinueGrant() (redirect *Redirect, err error) {
	// Read payment meta data written to file by NewRedirect
	b, err := fileutil.ReadAll(paymentMetaData)
	if err != nil {
		return redirect, err
	}
	err = json.Unmarshal(b, &redirect)
	if err != nil {
		return redirect, errors.WithStack(err)
	}

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
	err = fileutil.WriteBytes(paymentMetaData, out)
	if err != nil {
		return redirect, errors.WithStack(err)
	}

	return redirect, nil
}
