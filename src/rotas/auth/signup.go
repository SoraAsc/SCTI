package auth

import (
  "SCTI/fileserver"
  DB "SCTI/database"
  "encoding/json"
  "net/http"
  "log"
  "fmt"
)

func (h *Handler) GetSignup(w http.ResponseWriter, r *http.Request) {
  var t = fileserver.Execute("template/signup.gohtml")
  t.Execute(w, nil)
}

func (h *Handler) PostSignup(w http.ResponseWriter, r *http.Request) {
  var user User
  if r.Header.Get("Content-type") == "application/json" {
    err := json.NewDecoder(r.Body).Decode(&user)
    if err != nil {
      log.Fatal(err)
    }
  } else {
    if err := r.ParseForm(); err != nil {
      fmt.Println("r.Form dentro if: ", r.Form)
      log.Fatal(err)
    }
    user.Email = r.FormValue("Email")
    user.Password = r.FormValue("Senha")
  }

  found, err := DB.UserExists(user.Email)
  if err != nil {
    fmt.Println("DB UserExists Failed")
    return
  }

  if found {
    fmt.Println("User already exists")
    return
  }

  hash, _ := HashPassword(user.Password)
  err = DB.CreateUser(user.Email, hash)
  if err != nil {
    fmt.Println("Creating user in DB failed")
    return
  }
}

