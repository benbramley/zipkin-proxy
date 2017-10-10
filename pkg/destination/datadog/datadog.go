package datadog

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"sync"
	"time"

	"github.com/Sirupsen/logrus"
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

	// ContentType is the HTTP content type when sending data to the DataDog agent
	ContentType = "application/json"
)

// Shipper is a DataDog destination for sending openzipkin spans to
type Shipper struct {
	mu sync.Mutex

	agentHost string

	// configurable options
	flushDuration time.Duration
	bufferSize    int

	spans []*zipkin.Span
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
func (s *Shipper) SendSpan(span *zipkin.Span) error {
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
	s.mu.Lock()
	defer s.mu.Unlock()

	if len(s.spans) == 0 {
		logrus.Debug("skipping send, no spans present")
		return
	}

	logrus.WithField("length", len(s.spans)).Debug("sending spans")

	spans := make([]*Span, len(s.spans))

	for _, zs := range s.spans {
		spans = append(spans, SpanFromZipkin(zs))
	}

	req := [][]*Span{spans}
	buf := new(bytes.Buffer)
	if err := json.NewEncoder(buf).Encode(req); err != nil {
		logrus.WithError(err).Error("could not marshal JSON for DataDog")
	}

	agentReq, err := http.NewRequest("PUT", fmt.Sprintf("http://%s/%s/traces"), buf)
	if err != nil {
		logrus.WithError(err).Error("could not generate request for datadog")
	}

	agentReq.Header.Set("Content-Type", ContentType)

	_, err = http.DefaultClient.Do(agentReq)
	if err != nil {
		logrus.WithError(err).Error("failed sending request to the datadog agent")
	}
}
