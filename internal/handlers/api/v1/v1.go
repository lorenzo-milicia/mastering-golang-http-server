package v1

import (
	"github.com/gorilla/mux"
	"go.lorenzomilicia.com/go-master/http-server/internal/services/greet"
)

type Handler struct {
	GreetService *greet.Service
}

func (h *Handler) Route(router *mux.Router) {
	router.Handle("/greet", handleGreet(h.GreetService)).Methods("GET")
}
