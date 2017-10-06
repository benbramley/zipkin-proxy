package logs

import (
	"encoding/json"
	"fmt"

	"github.com/namely/zipkin-proxy/pkg/destination"
	"github.com/namely/zipkin-proxy/pkg/zipkin"
)

type LogShipper struct{}

var _ destination.Interface = &LogShipper{}

func (ls *LogShipper) SendSpan(s zipkin.Span) error {
	b, err := json.Marshal(s)
	if err != nil {
		return err
	}

	fmt.Println(string(b))
	return nil
}
