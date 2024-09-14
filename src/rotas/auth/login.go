package auth

import (
	DB "SCTI/database"
	Erros "SCTI/erros"
	"SCTI/fileserver"
	HTMX "SCTI/htmx"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

func GetLogin(w http.ResponseWriter, r *http.Request) {
	var t = fileserver.Execute("template/login.gohtml")
	t.Execute(w, nil)
}

func PostLogin(w http.ResponseWriter, r *http.Request) {
	var user User
	if r.Header.Get("Content-type") == "application/json" {
		err := json.NewDecoder(r.Body).Decode(&user)
		if err != nil {
			HTMX.Failure(w, "Falha no login: ", err)
			return
		}
	} else {
		if err := r.ParseForm(); err != nil {
			HTMX.Failure(w, "Falha no login: ", err)
			return
		}
		user.Email = r.FormValue("Email")
		user.Password = r.FormValue("Senha")
	}

	if user.Email == "" {
		HTMX.Failure(w, "Erro ao entrar: ", fmt.Errorf("Campo do email está vazio"))
		return
	}

	if user.Password == "" {
		HTMX.Failure(w, "Erro ao entrar: ", fmt.Errorf("Campo da senha está vazio"))
		return
	}

	login, uuid, err := VerifyLogin(user, w)

	if err != nil {
		HTMX.Failure(w, "Falha no login: ", err)
		return
	}

	if !login {
		HTMX.Failure(w, "Falha no login: ", fmt.Errorf("Erro desconhecido"))
		return
	}

	cookie := http.Cookie{
		Name:     "accessToken",
		Value:    uuid,
		Expires:  time.Now().Add(2 * 24 * time.Hour),
		Secure:   false,
		HttpOnly: true,
		Path:     "/",
		SameSite: http.SameSiteLaxMode,
	}

	http.SetCookie(w, &cookie)
	_ = VerifyAdmin(w, r, uuid)

	if r.Header.Get("HX-Request") == "true" {
		w.Header().Set("HX-Redirect", "/dashboard")
		w.WriteHeader(http.StatusOK)
		return
	}

	http.Redirect(w, r, "/dashboard", http.StatusSeeOther)
}

func VerifyLogin(user User, w http.ResponseWriter) (login bool, uuid string, err error) {
	found, err := DB.UserExists(user.Email)

	if err != nil {
		return false, "", fmt.Errorf("DB check query Failed")
	}

	if !found {
		return false, "", fmt.Errorf("User not found")
	}

	uuid = fmt.Sprint(DB.GetUUID(user.Email))
	if CheckPasswordHash(user.Password, DB.GetHash(user.Email)) && uuid != "" {
		return true, uuid, nil
	}
	return false, "", fmt.Errorf("Senha inválida")
}

func GetLogoff(w http.ResponseWriter, r *http.Request) {
	http.SetCookie(w, &http.Cookie{
		Name:     "accessToken",
		Value:    "",
		MaxAge:   -1,
		Secure:   false,
		HttpOnly: true,
		Path:     "/",
		SameSite: http.SameSiteLaxMode,
	})
	http.SetCookie(w, &http.Cookie{
		Name:     "Admin",
		Value:    "",
		MaxAge:   -1,
		Secure:   false,
		HttpOnly: true,
		Path:     "/",
		SameSite: http.SameSiteLaxMode,
	})
	http.Redirect(w, r, "/login", http.StatusSeeOther)
}

func VerifyAdmin(w http.ResponseWriter, r *http.Request, uuid string) bool {
	admcookie, err := r.Cookie("Admin")
	if err != nil {
		if err != http.ErrNoCookie {
			Erros.LogError("auth/login", fmt.Errorf("Error reading admin cookie: %v", err.Error()))
			Erros.HttpError(w, "auth/login", fmt.Errorf("Error reading admin cookie: %v", err.Error()))
			return false
		}
		fmt.Println("login.go: Admin cookie not present")
	} else {
		if admcookie.Value == uuid {
			return true
		} else {
			http.SetCookie(w, &http.Cookie{
				Name:     "Admin",
				Value:    "",
				MaxAge:   -1,
				Secure:   false,
				HttpOnly: true,
				Path:     "/",
				SameSite: http.SameSiteLaxMode,
			})
		}
	}

	isAdmin := DB.GetAdmin(uuid)

	if isAdmin {
		admcookie := http.Cookie{
			Name:     "Admin",
			Value:    uuid,
			Expires:  time.Now().Add(2 * 24 * time.Hour),
			Secure:   false,
			HttpOnly: true,
			Path:     "/",
			SameSite: http.SameSiteLaxMode,
		}
		http.SetCookie(w, &admcookie)
		return true
	}

	return false
}
