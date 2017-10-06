package zipkin

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestZipkinTraceIDSCanConvertToInts(t *testing.T) {
	span := Span{TraceID: "1c8abcbdbabdbc45"}

	i, err := span.TraceIDInt()
	require.NoError(t, err)
	assert.Equal(t, uint64(2056663702915890245), i)
}
