package ih2024

import (
	"github.com/shopd/shopd/go/plugin/ih2024/config"
	ps "github.com/shopd/shopd/go/services/plugin"
)

type Handler struct {
	// conf is the plugin config,
	// the domain config is available on s.DomainQX.DomainConf,
	// and the shopd app config is available on s.DomainQX.Conf
	conf    *config.Config
	s       *ps.Services
	enabled bool
}

// Name is used as the top level subject for topics related to this extension.
// See other top level subjects in msg package. May also be used as
// path prefix of the topic named param for webhook routes etc
func (ph *Handler) Name() string {
	return "ih2024"
}

func (ph *Handler) Info() string {
	return ""
}

func (ph *Handler) Enabled() bool {
	return ph.enabled
}

func (ph *Handler) Enable() (err error) {
	// Register publishers and subscribers
	err = ph.registerPubSub()
	if err != nil {
		return err
	}

	ph.enabled = true
	return nil
}

func (ph *Handler) Disable() error {
	ph.enabled = false
	return nil
}

func New(conf *config.Config, ps *ps.Services) *Handler {
	return &Handler{
		conf:    conf,
		s:       ps,
		enabled: false,
	}
}
