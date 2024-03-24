package main

import (
	"SCTI/fileserver"
	"SCTI/middleware"
	"html/template"
	"log"
	"net/http"
)

func main() {
	fileserver.RunFileServer()

	mux := http.NewServeMux()

	mux.HandleFunc("GET /", HomeHandler)

	server := http.Server{
		Addr:    ":8080",
		Handler: middleware.EndpointLogging(mux),
	}

	log.Fatal(server.ListenAndServe())
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	fileserver.T = template.Must(template.ParseFiles("template/index.gohtml"))
	fileserver.T.Execute(w, nil)
}
