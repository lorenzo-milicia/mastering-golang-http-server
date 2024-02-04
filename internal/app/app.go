package app

import (
	v1 "go.lorenzomilicia.com/go-master/http-server/internal/handlers/api/v1"
	"go.lorenzomilicia.com/go-master/http-server/internal/handlers/private"
	"go.lorenzomilicia.com/go-master/http-server/internal/inmemory"
	"go.lorenzomilicia.com/go-master/http-server/internal/services/greet"
	"go.lorenzomilicia.com/go-master/http-server/internal/services/project"
	"net/http"
)

type Server struct {
	Router *http.ServeMux
}

func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.Router.ServeHTTP(w, r)
}

func (s *Server) Routes() {
	// s.Router.Use(LoggerMiddleware)
	v1handler := v1.Handler{
		GreetService: &greet.Service{},
		ProjectService: &project.Service{
			Repository: &inmemory.InMemoryProjectRepository{},
		},
	}
	internalHandler := private.Handler{}
	v1handler.RegisterHandlers("/api/v1", s.Router)
	internalHandler.RegisterHandlers("/internal", s.Router)
}
