package auth

import (
	DB "SCTI/database"
	Erros "SCTI/erros"
	HTMX "SCTI/htmx"
	"fmt"
	"html/template"
	"net/http"
	"net/url"
)

type TrocarData struct {
	Email string
}

func GetTrocar(w http.ResponseWriter, r *http.Request) {
	code := r.URL.Query().Get("code")
	encodedEmail := r.URL.Query().Get("email")

	if code == "" || encodedEmail == "" {
		Erros.HttpError(w, "auth/verify", fmt.Errorf("Code or Email parameter is missing"))
		return
	}

	email, err := url.QueryUnescape(encodedEmail)
	if err != nil {
		Erros.HttpError(w, "auth/verify", fmt.Errorf("Invalid Email format"))
		return
	}

	DB_Code, err := DB.GetCodeByEmail(email)
	if err != nil {
		Erros.HttpError(w, "auth/verify", fmt.Errorf("Database error"))
		return
	}

	if DB_Code == "" {
		Erros.HttpError(w, "auth/verify", fmt.Errorf("User didn't have a verification code"))
		return
	}

	if DB_Code != code {
		Erros.HttpError(w, "auth/verify", fmt.Errorf("Invalid verification code"))
		return
	}

	data := TrocarData{
		Email: email,
	}

	tmpl, err := template.ParseFiles("template/trocar.gohtml")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	tmpl.ExecuteTemplate(w, "trocar", data)
}

func PostTrocar(w http.ResponseWriter, r *http.Request) {
	email := r.FormValue("Email")
	senha := r.FormValue("Senha")

	uuid := DB.GetUUID(email)
	fmt.Println(uuid)

	err := DB.ChangeUserPassword(uuid, senha)
	if err != nil {
		HTMX.Failure(w, "Não foi possível trocar a senha: ", err)
		return
	}
	HTMX.Success(w, "Senha trocada com sucesso")
}
