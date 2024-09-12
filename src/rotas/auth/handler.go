package auth

import (
	"net/http"

	"golang.org/x/crypto/bcrypt"
)

type User struct {
	Email    string
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
	mux.HandleFunc("GET /login", GetLogin)
	mux.HandleFunc("POST /login", PostLogin)
	mux.HandleFunc("GET /signup", GetSignup)
	mux.HandleFunc("POST /signup", PostSignup)
	mux.HandleFunc("GET /logoff", GetLogoff)
	mux.HandleFunc("GET /verify", GetVerify)
	mux.HandleFunc("GET /delete", GetDelete)
	mux.HandleFunc("GET /senha", GetSenha)
	mux.HandleFunc("POST /senha", PostSenha)
	mux.HandleFunc("GET /trocar", GetTrocar)
	mux.HandleFunc("POST /trocar", PostTrocar)
}
