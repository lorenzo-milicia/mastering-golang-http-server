package main

import (
	"fmt"
	"net/http"
	"os"

	"go.lorenzomilicia.com/go-master/http-server/internal/app"
)

func main() {
	server := &app.Server{
		Router: http.NewServeMux(),
	}
	server.Routes()

	err := http.ListenAndServe(":8080", server)
	if err != nil {
		fmt.Printf("error: %v", err)
		os.Exit(1)
	}
}
