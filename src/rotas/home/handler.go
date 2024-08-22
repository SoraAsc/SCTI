package home

import (
	DB "SCTI/database"
	"SCTI/rotas/notfound"
	"fmt"
	"html/template"
	"net/http"
)

type HomeData struct {
  Activities []DB.Activity
}

func GetHome(w http.ResponseWriter, r *http.Request) {
  //404 page handler
  if r.URL.Path != "/"{
    notfound.NotFound(w, r)
    return
  }

  a, err := DB.GetAllActivities()
  if err != nil {
    fmt.Println("Couldn't get Activities")
  }

  data := HomeData {
    Activities: a,
  }

  tmpl, err := template.ParseFiles("template/index.gohtml")
  if err != nil {
    http.Error(w, err.Error(), http.StatusInternalServerError)
    return
  }
  tmpl.ExecuteTemplate(w, "index", data)
}

func RegisterRoutes(mux *http.ServeMux) {
  mux.HandleFunc("/", GetHome)
}
