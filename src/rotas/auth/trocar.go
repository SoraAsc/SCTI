package auth

import (
  "fmt"
  "html/template"
  "net/http"
  "net/url"
  DB "SCTI/database"
  HTMX "SCTI/htmx"
)

type TrocarData struct {
  Email string
}

func GetTrocar(w http.ResponseWriter, r *http.Request) {
  email, err := url.QueryUnescape(r.URL.Query().Get("email"))

  fmt.Println(email)

  if err != nil {
    return
  }

  data := TrocarData{
    Email: email,
  }

  tmpl, err := template.ParseFiles("template/trocar.gohtml")
  if err != nil {
    http.Error(w, err.Error(), http.StatusInternalServerError)
    return
  }
  tmpl.ExecuteTemplate(w, "trocar", data)
}

func PostTrocar(w http.ResponseWriter, r *http.Request) {
  email := r.FormValue("Email")
  senha := r.FormValue("Senha")

  uuid := DB.GetUUID(email)
  fmt.Println(uuid)

  err := DB.ChangeUserPassword(uuid, senha)
  if err != nil {
    HTMX.Failure(w, "Não foi possível trocar a senha: ", err)
    return
  }
  HTMX.Success(w, "Senha trocada com sucesso")
}
