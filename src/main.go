package main

import (
	// "SCTI/about"
	"SCTI/fileserver"
	"SCTI/middleware"
	"html/template"
	"log"
	"net/http"
)

func main() {
	fileserver.RunFileServer()

	mux := http.NewServeMux()

	mux.HandleFunc("GET /", Home)

	server := http.Server{
		Addr:    ":8080",
		Handler: middleware.RouteLogging(mux),
	}

	log.Fatal(server.ListenAndServe())
}

func Home(w http.ResponseWriter, r *http.Request) {
	fileserver.T = template.Must(template.ParseFiles("template/index.gohtml", "template/footer.html"))
	fileserver.T.Execute(w, nil)
}
