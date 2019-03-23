package httpserver

import (
	"net/http"

	"github.com/moonrhythm/parapet"
)

// Server is the http server
type Server struct {
	Addr    string
	Handler http.Handler
}

// ListenAndServe listens and serve web server
func (s Server) ListenAndServe() error {
	svc := parapet.NewBackend()
	svc.Addr = s.Addr
	svc.Handler = s.Handler
	return svc.ListenAndServe()
}
