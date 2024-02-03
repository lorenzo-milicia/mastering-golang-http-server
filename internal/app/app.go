package app

import (
	"net/http"
	"os"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	v1 "go.lorenzomilicia.com/go-master/http-server/internal/handlers/api/v1"
	"go.lorenzomilicia.com/go-master/http-server/internal/handlers/private"
	"go.lorenzomilicia.com/go-master/http-server/internal/inmemory"
	"go.lorenzomilicia.com/go-master/http-server/internal/services/greet"
	"go.lorenzomilicia.com/go-master/http-server/internal/services/project"
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
		ProjectService: &project.Service{
			Repository: &inmemory.InMemoryProjectRepository{},
		},
	}
	internalHandler := private.Handler{}
	routes := map[string]Router{
		"/api/v1":   &v1handler,
		"/internal": &internalHandler,
	}
	for path, router := range routes {
		subr := s.Router.PathPrefix(path).Subrouter()
		router.Route(subr)
	}
}

func LoggerMiddleware(h http.Handler) http.Handler {
	return handlers.LoggingHandler(os.Stdout, h)
}

type Router interface {
	Route(*mux.Router)
}
