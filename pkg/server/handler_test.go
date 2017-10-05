package server

import (
	"bytes"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/namely/zipkin-proxy/pkg/zipkin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

type MockDestination struct {
	spans    []zipkin.Span
	errorOut bool
}

var fakeSpans = []byte(`[{"traceId":"string","name":"string","parentId":"string","id":"string","timestamp":0,"duration":0,"debug":true,"annotations":[{"endpoint":{"serviceName":"string","ipv4":"string","ipv6":"string","port":0},"timestamp":0,"value":"string"}],"binaryAnnotations":[{"key":"string","value":"string","endpoint":{"serviceName":"string","ipv4":"string","ipv6":"string","port":0}}]}]`)

func (md *MockDestination) SendSpan(s zipkin.Span) error {
	if md.errorOut {
		return errors.New("ooga booga!")
	}

	md.spans = append(md.spans, s)
	return nil
}

func TestServeHTTPSendsSpans(t *testing.T) {
	t.Run("Successful shipping of spans", func(t *testing.T) {
		d := &MockDestination{}
		h := &handler{d}

		svr := httptest.NewServer(h)
		body := bytes.NewBuffer(fakeSpans)

		resp, err := http.Post(svr.URL, "application/json", body)
		require.NoError(t, err)
		assert.Equal(t, resp.StatusCode, http.StatusAccepted)

		assert.Len(t, d.spans, 1, "fake destination has the span appended to it")
	})

	t.Run("The server errors on bad JSON", func(t *testing.T) {
		d := &MockDestination{}
		h := &handler{d}

		svr := httptest.NewServer(h)

		resp, err := http.Post(svr.URL, "application/json", nil)
		require.NoError(t, err)
		assert.Equal(t, resp.StatusCode, http.StatusBadRequest)
	})

	t.Run("The destination shipper errors out changing the status code to insufficient storage", func(t *testing.T) {
		d := &MockDestination{}
		h := &handler{d}

		svr := httptest.NewServer(h)
		body := bytes.NewBuffer(fakeSpans)
		d.errorOut = true

		resp, err := http.Post(svr.URL, "application/json", body)
		require.NoError(t, err)
		assert.Equal(t, resp.StatusCode, http.StatusInsufficientStorage)
	})
}
