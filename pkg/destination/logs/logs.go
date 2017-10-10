package logs

import (
	"encoding/json"
	"fmt"

	"github.com/namely/zipkin-proxy/pkg/destination"
	"github.com/namely/zipkin-proxy/pkg/zipkin"
)

// LogShipper implements the destination.Interface and just sends all incoming
// spans to STDOUT in JSON format.
type LogShipper struct{}

var _ destination.Interface = &LogShipper{}

// SendSpan implements destination.Interface
// All spans are marshalled as JSON and sent to Stdout
func (ls *LogShipper) SendSpan(s *zipkin.Span) error {
	b, err := json.Marshal(s)
	if err != nil {
		return err
	}

	fmt.Println(string(b))
	return nil
}
