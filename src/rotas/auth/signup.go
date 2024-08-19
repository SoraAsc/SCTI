package auth

import (
  DB "SCTI/database"
  "SCTI/fileserver"
  "encoding/json"
  "fmt"
  "log"
  "net/http"

  "github.com/google/uuid"
)

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
      log.Fatal(err)
    }
  } else {
    if err := r.ParseForm(); err != nil {
      fmt.Println("r.Form dentro if: ", r.Form)
      log.Fatal(err)
    }
    name = r.FormValue("Nome")
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


  UUID := uuid.New()
  UUIDString := UUID.String()

  hash, _ := HashPassword(user.Password)
  err = DB.CreateUser(user.Email, hash, UUIDString, name)
  if err != nil {
    fmt.Printf("Creating user in DB failed: %v\n", err)
    return
  }


  http.Redirect(w, r, "/login", http.StatusSeeOther)
}
