package dashboard

import (
  DB "SCTI/database"
  HTMX "SCTI/htmx"
  "fmt"
  "net/http"
  "strconv"
)

func CheckAdmin(w http.ResponseWriter, r *http.Request) bool {
  admcookie, err := r.Cookie("Admin")
  if err != nil {
    HTMX.Failure(w, "Falha ao ler cookie de Admin: ", err)
    return false
  }

  logincookie, err := r.Cookie("accessToken")
  if err != nil {
    HTMX.Failure(w, "Falha ao ler cookie de Login: ", err)
    return false
  }

  if admcookie.Value != logincookie.Value {
    HTMX.Failure(w, "Admin inválido: ", fmt.Errorf("Cookie de Login e Admin diferem no usuário"))
    http.SetCookie(w, &http.Cookie{
      Name:   "Admin",
      Value:  "",
      MaxAge: -1,
      Secure: false,
      HttpOnly: true,
      Path: "/",
      SameSite: http.SameSiteLaxMode,
    })
    return false
  }
  return true
}

func SetAdmin(w http.ResponseWriter, r *http.Request) {
  if !CheckAdmin(w, r) {
    HTMX.Failure(w, "Endpoint exclusivo de admins", fmt.Errorf("Acesso proibido a usuários não admin"))
    return
  }
  email := r.FormValue("Email")

  err := DB.SetAdmin(DB.GetUUID(email), true)
  if err != nil {
    HTMX.Failure(w, "Falha ao criar o Admin: ", err)
    return
  }

  HTMX.Success(w, "Admin criado com sucesso")
}

func PostActivity(w http.ResponseWriter, r* http.Request) {
  if !CheckAdmin(w, r) {
    HTMX.Failure(w, "Endpoint exclusivo de admins", fmt.Errorf("Acesso proibido a usuários não admin"))
    return
  }

  var a DB.Activity
  a.Spots, _ = strconv.Atoi(r.FormValue("spots"))
  a.Activity_type = r.FormValue("type")
  a.Room = r.FormValue("room")
  a.Speaker = r.FormValue("speaker")
  a.Topic = r.FormValue("topic")
  a.Description = r.FormValue("description")
  a.Time = r.FormValue("time")
  a.Day, _ = strconv.Atoi(r.FormValue("day"))

  _, err := DB.CreateActivity(a)
  if err != nil {
    HTMX.Failure(w, "Falha ao criar atividade", err)
    return
  }
  HTMX.Success(w, "Atividade criada com sucesso")
}
