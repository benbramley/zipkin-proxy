package datadog

import (
	"sync"
	"time"

	"github.com/namely/zipkin-proxy/pkg/destination"
	"github.com/namely/zipkin-proxy/pkg/zipkin"
)

const (
	// APIVersion defines which datadog agent version we're using for sending
	// traces
	APIVersion = "v0.3"

	// TracesPath is the path after the datadog version in the URL for us to send
	// our JSON posts to
	TracesPath = "/traces"
)

// Shipper is a DataDog destination for sending openzipkin spans to
type Shipper struct {
	mu sync.Mutex

	agentHost string

	// configurable options
	flushDuration time.Duration
	bufferSize    int

	spans zipkin.Spans
}

var _ destination.Interface = &Shipper{}

// ShipperOpt configures a shipper instance
type ShipperOpt func(s *Shipper)

// NewShipper initializes a new shipper destination
func NewShipper(agentHost string, opts ...ShipperOpt) *Shipper {
	s := &Shipper{
		agentHost:     agentHost,
		flushDuration: time.Second * 3,
	}

	for _, opt := range opts {
		opt(s)
	}

	return s
}

// SendSpan implements destination.Interface by sending spans to datadog's
// tracing APM
func (s *Shipper) SendSpan(span zipkin.Span) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.spans = append(s.spans, span)

	return nil
}

func (s *Shipper) work() {
	ticker := time.NewTicker(s.flushDuration)

	for {
		<-ticker.C
		s.shipSpans()
	}
}

func (s *Shipper) shipSpans() {

}
