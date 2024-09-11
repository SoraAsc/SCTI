package auth

import (
	DB "SCTI/database"
	"SCTI/fileserver"
	HTMX "SCTI/htmx"
	"encoding/json"
	"fmt"
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
			HTMX.Failure(w, "Falha no signup: ", err)
			return
		}
	} else {
		if err := r.ParseForm(); err != nil {
			HTMX.Failure(w, "Falha no signup: ", err)
			return
		}
		name = r.FormValue("Nome")
		user.Email = r.FormValue("Email")
		user.Password = r.FormValue("Senha")
	}

	if name == "" {
		HTMX.Failure(w, "Erro ao cadastrar: ", fmt.Errorf("Campo do nome está vazio"))
		return
	}

	if user.Email == "" {
		HTMX.Failure(w, "Erro ao cadastrar: ", fmt.Errorf("Campo do email está vazio"))
		return
	}

	if user.Password == "" {
		HTMX.Failure(w, "Erro ao cadastrar: ", fmt.Errorf("Campo da senha está vazio"))
		return
	}

	found, err := DB.UserExists(user.Email)
	if err != nil {
		HTMX.Failure(w, "Falha no signup: ", err)
		return
	}

	if found {
		HTMX.Failure(w, "Falha no signup: ", fmt.Errorf("Usuário já existe"))
		return
	}

	UUID := uuid.New()
	UUIDString := UUID.String()

	hash, _ := HashPassword(user.Password)
	err = DB.CreateUser(user.Email, hash, UUIDString, name)
	if err != nil {
		HTMX.Failure(w, "Falha no signup: ", err)
		return
	}

	if r.Header.Get("HX-Request") == "true" {
		w.Header().Set("HX-Redirect", "/login")
		w.WriteHeader(http.StatusOK)
		return
	}

	http.Redirect(w, r, "/login", http.StatusSeeOther)
}
