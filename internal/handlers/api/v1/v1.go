package v1

import (
	"github.com/gorilla/mux"
	"go.lorenzomilicia.com/go-master/http-server/internal/services/greet"
	"go.lorenzomilicia.com/go-master/http-server/internal/services/project"
)

type Handler struct {
	GreetService   *greet.Service
	ProjectService *project.Service
}

func (h *Handler) Route(router *mux.Router) {
	router.Handle("/greet", handleGreet(h.GreetService)).Methods("GET")
	projectRouter := router.PathPrefix("/project").Subrouter()
	handleProject(projectRouter, h.ProjectService)
}
