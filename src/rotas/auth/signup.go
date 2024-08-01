package auth

import (
  "SCTI/fileserver"
  "encoding/json"
  "net/http"
  "log"
  "fmt"
  "os"

  "github.com/google/uuid"
)

func (h *Handler) GetSignup(w http.ResponseWriter, r *http.Request) {
  var t = fileserver.Execute("template/signup.gohtml")
  t.Execute(w, nil)
}

func (h *Handler) PostSignup(w http.ResponseWriter, r *http.Request) {
  println("In PostSignup")

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

  if UserExists(user.Email) {
    println("User already exists")
    return
  }

  hash, _ := HashPassword(user.Password)

  file, err := os.OpenFile("passwords.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0600)
  if err != nil {
    log.Fatal(err)
  }
  defer file.Close()

  _, err = file.WriteString(fmt.Sprintf("%s:%s:%s\n", user.Email, hash, uuid.NewString()))
  if err != nil {
    log.Fatal(err)
    return
  }

  fmt.Println("E-Mail: ", user.Email)
  fmt.Println("Password: ", user.Password)
  fmt.Println("Hash: ", hash)
}

