package main

import (
	"net/http"

	"github.com/gorilla/mux"
	"go.lorenzomilicia.com/go-master/http-server/internal/app"
)

func main() {
	server := &app.Server{
		Router: mux.NewRouter(),
	}
	server.Routes()
	http.ListenAndServe(":8080", server)
}
