package datadog

import "github.com/namely/zipkin-proxy/pkg/zipkin"

// Spans is a collection of spans for datadog requests
type Spans [][]Span

// Span represents a span that datadog understands
type Span struct {
	Start    int64             `json:"start"`
	Duration int64             `json:"duration"`
	TraceID  uint64            `json:"trace_id"`
	SpanID   uint64            `json:"span_id"`
	ParentID uint64            `json:"parent_id"`
	Name     string            `json:"name"`
	Resource string            `json:"resource"`
	Service  string            `json:"service"`
	Type     string            `json:"type"`
	Error    int32             `json:"error"`
	Meta     map[string]string `json:"meta,omitempty"`
}

func SpanFromZipkin(zs zipkin.Span) (*Span, error) {
	s := &Span{
		Start:    zs.Timestamp,
		Duration: zs.Duration,
		Name:     zs.Name,
	}

	var err error
	s.TraceID, err = zs.TraceIDInt()

	return s, err
}
