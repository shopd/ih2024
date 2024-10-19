package ih2024

import (
	"github.com/mozey/ft"
)

// TODO Review this once callback payload is known
const GrantAccepted = "Accepted"
const GrantRejected = "Rejected"

type Event struct {
	Result ft.NString `json:"type"`
	Amount ft.NInt    `json:"amount"`
}

// Webhook wraps the callback event
type Webhook struct {
	Event
	OrderNo string
}
