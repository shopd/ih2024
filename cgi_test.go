package ih2024_test

import (
	"testing"

	"github.com/rs/zerolog/log"
	"github.com/shopd/shopd/go/plugin/ih2024"
	"github.com/shopd/shopd/go/plugin/ih2024/config"
	"github.com/shopd/shopd/go/share"
	"github.com/shopd/shopd/go/testutil"
)

// Run this test from shopd APP_DIR like this:
//
//	go clean -testcache && gotest -v ./go/plugin/ih2024/... -run TestNewRedirect
func TestNewRedirect(t *testing.T) {
	is, _, ps := testutil.SetupPluginServices(t)

	pConf := config.New()
	ph := ih2024.New(pConf, ps)

	nonce, err := share.GenerateToken(share.GenerateRandBytes(20))
	is.NoErr(err)
	redirect, err := ph.NewRedirect(ih2024.NewRedirectParams{
		// SuccessURL is left empty,
		// the ASE will redirect to a page that says ACCEPT or REJECT
		SuccessURL: "",
		Nonce:      nonce,
		Amount:     1000,
		OrderID:    "2nfp6UcSPaCwrqPYfq3OMPEBTTZ",
		OrderNo:    "000010",
	})
	is.NoErr(err)

	log.Info().Interface("redirect", redirect).Msg("")
}

// Run this test from shopd APP_DIR like this:
//
//	go clean -testcache && gotest -v ./go/plugin/ih2024/... -run TestContinueGrant
func TestContinueGrant(t *testing.T) {
	is, _, ps := testutil.SetupPluginServices(t)

	pConf := config.New()
	ph := ih2024.New(pConf, ps)

	redirect, err := ph.ContinueGrant()
	is.NoErr(err)

	log.Info().Interface("redirect", redirect).Msg("")
}
