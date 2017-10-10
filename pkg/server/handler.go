package server

import (
	"net/http"

	"git.apache.org/thrift.git/lib/go/thrift"

	"github.com/Sirupsen/logrus"
	"github.com/namely/zipkin-proxy/pkg/destination"
	"github.com/namely/zipkin-proxy/pkg/zipkin"
	"github.com/pkg/errors"
)

type handler struct {
	destination destination.Interface
}

// ServeHTTP handles accepting spans and shipping them to the configured
// destination
func (h *handler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	defer req.Body.Close()

	protocol := thrift.NewTBinaryProtocolTransport(thrift.NewStreamTransportR(req.Body))

	_, size, err := protocol.ReadListBegin()
	if err != nil {
		logrus.WithError(err).Error()
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if size <= 0 || size > 32*1024 {
		logrus.WithError(err).Error()
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	statusCode := http.StatusAccepted
	for idx := 0; idx < size; idx++ {
		span := new(zipkin.Span)
		if err := span.Read(protocol); err != nil {
			logrus.WithError(err).Error()
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		if err := h.destination.SendSpan(span); err != nil {
			err = errors.Wrap(err, "could not send span to destination")
			logrus.WithError(err).Error("error shipping span")
			statusCode = http.StatusInsufficientStorage
		}
	}

	w.WriteHeader(statusCode)

	// var spans zipkin.Spans
	// if err := json.NewDecoder(buf).Decode(&spans); err != nil {
	// 	logrus.WithError(err).Error()
	// 	w.WriteHeader(http.StatusBadRequest)
	// 	w.Write([]byte(http.StatusText(http.StatusBadRequest)))
	// 	return
	// }
	//
	// buf.
	//
	// statusCode := http.StatusAccepted
	// for _, span := range spans {
	// 	if err := h.destination.SendSpan(span); err != nil {
	// 		err = errors.Wrap(err, "could not send span to destination")
	// 		logrus.WithError(err).Error("error shipping span")
	//
	// 		statusCode = http.StatusInsufficientStorage
	// 	}
	// }
	//
	// w.WriteHeader(statusCode)
}
