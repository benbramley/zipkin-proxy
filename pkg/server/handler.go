package server

import (
	"encoding/json"
	"net/http"

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
	var spans zipkin.Spans
	if err := json.NewDecoder(req.Body).Decode(&spans); err != nil {
		logrus.WithError(err).Error()
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(http.StatusText(http.StatusBadRequest)))
		return
	}

	statusCode := http.StatusAccepted
	for _, span := range spans {
		if err := h.destination.SendSpan(span); err != nil {
			err = errors.Wrap(err, "could not send span to destination")
			logrus.WithError(err).Error("error shipping span")

			statusCode = http.StatusInsufficientStorage
		}
	}

	w.WriteHeader(statusCode)
}
