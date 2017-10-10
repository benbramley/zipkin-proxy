package datadog

import "github.com/namely/zipkin-proxy/pkg/zipkin"

// Span represents a span that datadog understands
type Span struct {
	Start    int64             `json:"start"`
	Duration int64             `json:"duration"`
	TraceID  int64             `json:"trace_id"`
	SpanID   int64             `json:"span_id"`
	ParentID int64             `json:"parent_id"`
	Name     string            `json:"name"`
	Resource string            `json:"resource"`
	Service  string            `json:"service"`
	Type     string            `json:"type"`
	Error    int32             `json:"error"`
	Meta     map[string]string `json:"meta,omitempty"`
}

func SpanFromZipkin(zs *zipkin.Span) *Span {
	s := &Span{
		Start:    zs.GetTimestamp(),
		Duration: zs.GetDuration(),
		Name:     zs.GetName(),
		TraceID:  zs.GetTraceID(),
		ParentID: zs.GetParentID(),
		SpanID:   zs.GetID(),
	}

	return s
}
