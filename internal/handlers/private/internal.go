package private

import (
	"net/http"

	"github.com/gorilla/mux"
)

type Handler struct{}

func (h *Handler) Route(r *mux.Router) {
	r.Handle("/health", h.handleInternalHealth())
}

func (h *Handler) RegisterHandlers(path string, mux *http.ServeMux) {
	mux.Handle(path+"/health", h.handleInternalHealth())
}

func (h *Handler) handleInternalHealth() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
	}
}
