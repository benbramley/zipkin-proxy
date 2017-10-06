package zipkin

import (
	"fmt"
	"strconv"
)

// Spans represents a collection of traces
type Spans []Span

// Span represents an OpenZipkin trace message
type Span struct {
	TraceID           string             `json:"traceId"`
	Name              string             `json:"name"`
	ParentID          string             `json:"parentId"`
	ID                string             `json:"id"`
	Timestamp         int64              `json:"timestamp"`
	Duration          int64              `json:"duration"`
	Debug             bool               `json:"debug"`
	Annotations       []Annotation       `json:"annotations"`
	BinaryAnnotations []BinaryAnnotation `json:"binaryAnnotations"`
}

// Returns the TraceID for the span as a uint64 representation
// by parsing the hex encoded number
func (s Span) TraceIDInt() (uint64, error) {
	return strconv.ParseUint(fmt.Sprintf("0x%s", s.TraceID), 0, 64)
}

// BinaryAnnotation represents an OpenZipkin binary annotation message
type BinaryAnnotation struct {
	Key      string   `json:"key"`
	Value    string   `json:"value"`
	Endpoint Endpoint `json:"endpoint"`
}

// Annotation represents an OpenZipkin annotation message
type Annotation struct {
	Timestamp int      `json:"timestamp"`
	Value     string   `json:"value"`
	Endpoint  Endpoint `json:"endpoint"`
}

// Endpoint represents an OpenZipkin endpoint message
type Endpoint struct {
	ServiceName string `json:"serviceName"`
	Ipv4        string `json:"ipv4"`
	Ipv6        string `json:"ipv6"`
	Port        int    `json:"port"`
}
