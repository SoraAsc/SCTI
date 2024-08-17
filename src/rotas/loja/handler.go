package loja

import (
  "SCTI/fileserver"

  "net/http"
)

func GetLoja(w http.ResponseWriter, r *http.Request) {
  auth, err := r.Cookie("accessToken")
  if err != nil {
    // fmt.Println("Error Getting cookie:", err)
    http.Redirect(w, r, "/login", http.StatusSeeOther)
    return
  }

  if auth.Value == "-1" {
    // fmt.Println("Invalid accessToken")
    http.Redirect(w, r, "/login", http.StatusSeeOther)
  }
  var t = fileserver.Execute("template/loja.gohtml")
  t.Execute(w, nil)
}

func RegisterRoutes(mux *http.ServeMux) {
  mux.HandleFunc("GET /loja", GetLoja)
}
