package server

import (
	"net/http"

	"github.com/gorilla/mux"

	"github.com/namely/zipkin-proxy/pkg/destination"
)

// Server implements the API v1 JSON endpoint of Zipkin described here
// http://zipkin.io/zipkin-api/#/default/post_spans
// When receiving spans it should ship them to the desired destination
type Server struct {
	listenOn string
	handler  *handler
}

// NewServer configures a server for receiving spans
func NewServer(listenOn string, dest destination.Interface) *Server {
	return &Server{
		listenOn: listenOn,
		handler:  &handler{dest},
	}
}

// Start starts the Zipkin proxy server for accepting connections
func (s *Server) Start() error {
	r := mux.NewRouter()
	r.Handle("/api/v1/spans", s.handler).Methods(http.MethodPost)

	return http.ListenAndServe(s.listenOn, r)
}
