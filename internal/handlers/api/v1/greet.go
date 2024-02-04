package v1

import (
	"encoding/json"
	"net/http"

	"go.lorenzomilicia.com/go-master/http-server/internal/services/greet"
)

func handleGreet22(path string, mux *http.ServeMux, s *greet.Service) {
	mux.Handle("GET "+path+"/{name}", handleGreet(s))
}

func handleGreet(s *greet.Service) http.HandlerFunc {
	type request struct {
		Name string
	}
	type response struct {
		Greeting string `json:"greeting"`
	}
	return func(w http.ResponseWriter, r *http.Request) {
		request := request{
			Name: r.PathValue("name"),
		}
		response := response{
			Greeting: s.Greet(request.Name),
		}
		body, err := json.Marshal(response)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		w.Write(body)
	}
}
