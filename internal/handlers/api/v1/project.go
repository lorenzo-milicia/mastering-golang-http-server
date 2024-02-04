package v1

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"go.lorenzomilicia.com/go-master/http-server/internal/model"
	"go.lorenzomilicia.com/go-master/http-server/internal/services/project"
)

func handleProject(r *mux.Router, s *project.Service) {
	r.Handle("", handleProjectGet(s)).Methods("GET")
	r.Handle("", handleProjectPost(s)).Methods("POST")
}

func handleProject22(path string, mux *http.ServeMux, s *project.Service) {
	mux.Handle("GET "+path+"/", handleProjectGet(s))
	mux.Handle("POST "+path+"/", handleProjectPost(s))
}

func handleProjectGet(s *project.Service) http.HandlerFunc {
	type project struct {
		Name string `json:"name"`
	}
	return func(w http.ResponseWriter, r *http.Request) {
		projects, err := s.GetAll()
		if err != nil {
			w.WriteHeader(500)
		}
		var response []project
		for _, p := range projects {
			response = append(response, project{Name: p.Name})
		}
		json, err := json.Marshal(response)
		w.Write(json)
	}
}
func handleProjectPost(s *project.Service) http.HandlerFunc {
	type request struct {
		Name string `json:"name"`
	}
	return func(w http.ResponseWriter, r *http.Request) {
		var req request
		json.NewDecoder(r.Body).Decode(&req)
		project := model.Project{
			Name: req.Name,
		}
		err := s.Save(project)
		if err != nil {
			w.WriteHeader(500)
		}
		w.WriteHeader(200)
	}
}
