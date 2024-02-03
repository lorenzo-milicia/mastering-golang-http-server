package v1

import (
	"encoding/json"
	"net/http"

	"go.lorenzomilicia.com/go-master/http-server/internal/services/greet"
)

func handleGreet(s *greet.Service) http.HandlerFunc {
	type request struct {
		Name string
	}
	type response struct {
		Greeting string `json:"greeting"`
	}
	return func(w http.ResponseWriter, r *http.Request) {
		request := request{
			Name: r.URL.Query().Get("name"),
		}
		response := response{
			Greeting: s.Greet(request.Name),
		}
		body, err := json.Marshal(response)
		if err != nil {
			w.WriteHeader(500)
		}
		w.Write(body)
	}
}
