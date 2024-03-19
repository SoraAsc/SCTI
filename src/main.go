package main

import (
	"log"
	"net/http"
  "html/template"
  "SCTI/about"
  "SCTI/fileserver"
)

func main() {
  fileserver.RunFileServer()

  http.HandleFunc("/", HomeHandler)
  http.HandleFunc("/home/", about.Endpoint)

  log.Fatal(http.ListenAndServe(":8080", nil))
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
  fileserver.T = template.Must(template.ParseFiles("template/index.gohtml"))
  fileserver.T.Execute(w, nil)
}
