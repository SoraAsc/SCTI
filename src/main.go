package main

import (
	"log"
	"net/http"
  "html/template"
)

var t *template.Template

func main() {
  fs := http.FileServer(http.Dir("./static"))
  http.Handle("/static/", http.StripPrefix("/static/", fs))

  http.HandleFunc("/", HomeHandler)

  log.Fatal(http.ListenAndServe(":8080", nil))
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
  t := template.Must(template.ParseFiles("template/index.gohtml"))
  t.Execute(w, nil)
}
