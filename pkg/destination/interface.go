package destination

import "github.com/namely/zipkin-proxy/pkg/zipkin"

// Interface defines a common interface for shipping spans to another destination
// such as datadog APM or whatever adapter you decide to implement
type Interface interface {
	// SendSpan should send the span to another APM monitoring service
	// like DataDog or similar
	SendSpan(span *zipkin.Span) error
}
