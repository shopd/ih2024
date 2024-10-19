package ih2024

type Redirect struct {
	RedirectURL string
	ContinueURI string
	AccessToken string
	QuoteID     string
	Amount      int64
}

type NewRedirectParams struct {
	InWalletAddressURL  string
	OutWalletAddressURL string
	KeyID               string
	PrivateKey          string
	SuccessURL          string
	Nonce               string
	Amount              int64
}

func (ph *Handler) NewRedirect(params NewRedirectParams) (redirect *Redirect, err error) {
	// TODO Run command `node cgi/redirect.js`
	return redirect, nil
}
