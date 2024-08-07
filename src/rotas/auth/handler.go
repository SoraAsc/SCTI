package auth

import (
  "fmt"
  "net/http"

  "golang.org/x/crypto/bcrypt"
)

type Handler struct{}

type User struct {
  Email string
  Password string 
}

func HashPassword(password string) (string, error) {
  bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
  return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
  err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
  return err == nil
}

func (h *Handler) GetPrintCookies(w http.ResponseWriter, req *http.Request) {
  var returnStr string
  for _, cookie := range req.Cookies() {
    returnStr = returnStr + cookie.Name + ":" + cookie.Value + "\n"
  }
  fmt.Fprint(w, returnStr)
}
