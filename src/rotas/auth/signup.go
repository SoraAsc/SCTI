package auth

import (
  DB "SCTI/database"
  "SCTI/fileserver"
  "encoding/json"
  "fmt"
  "net/http"

  "github.com/google/uuid"
)

func SignupFailed(w http.ResponseWriter, err error) {
  w.Header().Set("Content-Type", "text/html")
  w.Write([]byte(`
    <div class="failure">
    Falha no Signup:
    ` + err.Error() + `
    </div>
    `))
}

func GetSignup(w http.ResponseWriter, r *http.Request) {
  var t = fileserver.Execute("template/signup.gohtml")
  t.Execute(w, nil)
}

func PostSignup(w http.ResponseWriter, r *http.Request) {
  var user User
  var name string
  if r.Header.Get("Content-type") == "application/json" {
    err := json.NewDecoder(r.Body).Decode(&user)
    if err != nil {
      SignupFailed(w, err)
      return
    }
  } else {
    if err := r.ParseForm(); err != nil {
      SignupFailed(w, err)
      return
    }
    name = r.FormValue("Nome")
    user.Email = r.FormValue("Email")
    user.Password = r.FormValue("Senha")
  }

  found, err := DB.UserExists(user.Email)
  if err != nil {
    SignupFailed(w, err)
    return
  }

  if found {
    SignupFailed(w, fmt.Errorf("Usuário já existe"))
    return
  }

  UUID := uuid.New()
  UUIDString := UUID.String()

  hash, _ := HashPassword(user.Password)
  err = DB.CreateUser(user.Email, hash, UUIDString, name)
  if err != nil {
    SignupFailed(w, err)
    return
  }

  if r.Header.Get("HX-Request") == "true" {
    w.Header().Set("HX-Redirect", "/login")
    w.WriteHeader(http.StatusOK)
    return
  }

  http.Redirect(w, r, "/login", http.StatusSeeOther)
}
