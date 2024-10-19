package ih2024

import (
	"github.com/mozey/errors"
)

var ErrYoco = errors.NewCause("ih2024")

var ErrNotImplemented = errors.NewWithCause(ErrYoco, "not implemented")

// ErrOrderNo if orderNo matched none or more than one
var ErrOrderNo = func(orderNo string) error {
	return errors.NewWithCausef(ErrYoco, "no match for orderNo %s", orderNo)
}
