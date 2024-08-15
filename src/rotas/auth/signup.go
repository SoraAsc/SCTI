package auth

import (
	DB "SCTI/database"
	"SCTI/fileserver"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/smtp"
	"os"
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

  code := DB.GetCode(user.Email)
  if code == "" {
    fmt.Println("Error Getting the code")
    return
  }

  from := os.Getenv("GMAIL_SENDER")
  pass := os.Getenv("GMAIL_PASS")

  body := "Seu código de verificação é:\n" + code

  msg := "From: " + from + "\n" + "To: " + user.Email + "\n" + "Subject: Verificação de email SCTI\n\n" + body

  err = smtp.SendMail("smtp.gmail.com:587", smtp.PlainAuth("", from, pass, "smtp.gmail.com"), from, []string{user.Email}, []byte(msg))

  if err != nil {
    fmt.Printf("smtp error: %s\n", err)
    return
  }
}

