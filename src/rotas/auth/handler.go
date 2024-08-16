package auth

import (
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

func RegisterRoutes(mux *http.ServeMux) {
  handler := &Handler{}
  mux.HandleFunc("GET /login", handler.GetLogin)
  mux.HandleFunc("POST /login", handler.PostLogin)
  mux.HandleFunc("GET /signup", handler.GetSignup)
  mux.HandleFunc("POST /signup", handler.PostSignup)
  mux.HandleFunc("GET /logoff", handler.GetLogoff)
  mux.HandleFunc("GET /verify", handler.GetVerify)
}
