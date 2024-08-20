package dashboard

import (
  "encoding/json"
  "fmt"
  "log"
  "net/http"
)

func PostCadastros(w http.ResponseWriter, r *http.Request) {
  var id string
  if r.Header.Get("Content-type") == "application/json" {
    err := json.NewDecoder(r.Body).Decode(&id)
    if err != nil {
      log.Fatal(err)
    }
  } else {
    if err := r.ParseForm(); err != nil {
      fmt.Println("r.Form dentro if: ", r.Form)
      log.Fatal(err)
    }
    id = r.FormValue("ID")
  }
  fmt.Fprintf(w, "O ID é %s", id)
}

