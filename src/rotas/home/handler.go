package home

import (
  "SCTI/fileserver"
  "SCTI/rotas/notfound"
  "net/http"
)

func GetHome(w http.ResponseWriter, r *http.Request) {
  //404 page handler
  if r.URL.Path != "/"{
    notfound.NotFound(w, r)
    return
  }

  var t = fileserver.Execute("template/index.gohtml")
  t.Execute(w, nil)
}

func RegisterRoutes(mux *http.ServeMux) {
  mux.HandleFunc("/", GetHome)
}
