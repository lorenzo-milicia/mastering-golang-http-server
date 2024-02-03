package app

import (
	"net/http"
	"os"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	v1 "go.lorenzomilicia.com/go-master/http-server/internal/handlers/api/v1"
	"go.lorenzomilicia.com/go-master/http-server/internal/services/greet"
)

type Server struct {
	Router *mux.Router
}

func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.Router.ServeHTTP(w, r)
}

func (s *Server) Routes() {
	s.Router.Use(LoggerMiddleware)
	v1handler := v1.Handler{
		GreetService: &greet.Service{},
	}
	v1router := s.Router.PathPrefix("/api/v1").Subrouter()
	v1handler.Route(v1router)
}

func LoggerMiddleware(h http.Handler) http.Handler {
	return handlers.LoggingHandler(os.Stdout, h)
}
