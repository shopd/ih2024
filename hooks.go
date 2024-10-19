package ih2024

import (
	"net/url"

	"github.com/pkg/errors"
	"github.com/shopd/shopd/go/share"
)

// PaymentRedirect hook starts a new payment redirect workflow.
// User interaction is required to approve the payment.
// Webhooks from the payment processor is received by the message handler.
// The message handler completes the workflow.
// The order status page possibly polls for updated status
func (ph *Handler) PaymentRedirect(params share.PaymentRedirectParams) (
	redirectURL *url.URL, err error) {

	nonce, err := share.GenerateToken(share.GenerateRandBytes(20))
	if err != nil {
		return redirectURL, err
	}
	co, err := ph.NewRedirect(NewRedirectParams{
		InWalletAddressURL:  ph.conf.Ih2024InWalletAddressUrl(),
		OutWalletAddressURL: ph.conf.Ih2024OutWalletAddressUrl(),
		KeyID:               ph.conf.Ih2024KeyId(),
		PrivateKey:          ph.conf.Ih2024PrivateKey(),
		SuccessURL:          params.SuccessURL,
		Nonce:               nonce,
		Amount:              params.Order.Totals.Subtotal.Int64,
		// TODO For now just use wallet currency
		// Currency:   params.Currency,
	})
	if err != nil {
		return redirectURL, err
	}
	redirectURL, err = url.Parse(co.RedirectURL)
	if err != nil {
		return redirectURL, errors.WithStack(err)
	}

	return redirectURL, nil
}
