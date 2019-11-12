package performance

// Code generated by cdproto-gen. DO NOT EDIT.

import (
	"errors"

	"github.com/mailru/easyjson"
	"github.com/mailru/easyjson/jlexer"
	"github.com/mailru/easyjson/jwriter"
)

// Metric run-time execution metric.
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/Performance#type-Metric
type Metric struct {
	Name  string  `json:"name"`  // Metric name.
	Value float64 `json:"value"` // Metric value.
}

// SetTimeDomainTimeDomain time domain.
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/Performance#type-timeDomain
type SetTimeDomainTimeDomain string

// String returns the SetTimeDomainTimeDomain as string value.
func (t SetTimeDomainTimeDomain) String() string {
	return string(t)
}

// SetTimeDomainTimeDomain values.
const (
	SetTimeDomainTimeDomainTimeTicks   SetTimeDomainTimeDomain = "timeTicks"
	SetTimeDomainTimeDomainThreadTicks SetTimeDomainTimeDomain = "threadTicks"
)

// MarshalEasyJSON satisfies easyjson.Marshaler.
func (t SetTimeDomainTimeDomain) MarshalEasyJSON(out *jwriter.Writer) {
	out.String(string(t))
}

// MarshalJSON satisfies json.Marshaler.
func (t SetTimeDomainTimeDomain) MarshalJSON() ([]byte, error) {
	return easyjson.Marshal(t)
}

// UnmarshalEasyJSON satisfies easyjson.Unmarshaler.
func (t *SetTimeDomainTimeDomain) UnmarshalEasyJSON(in *jlexer.Lexer) {
	switch SetTimeDomainTimeDomain(in.String()) {
	case SetTimeDomainTimeDomainTimeTicks:
		*t = SetTimeDomainTimeDomainTimeTicks
	case SetTimeDomainTimeDomainThreadTicks:
		*t = SetTimeDomainTimeDomainThreadTicks

	default:
		in.AddError(errors.New("unknown SetTimeDomainTimeDomain value"))
	}
}

// UnmarshalJSON satisfies json.Unmarshaler.
func (t *SetTimeDomainTimeDomain) UnmarshalJSON(buf []byte) error {
	return easyjson.Unmarshal(buf, t)
}
