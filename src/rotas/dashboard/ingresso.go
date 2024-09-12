package dashboard

import (
	DB "SCTI/database"
	Erros "SCTI/erros"
	"SCTI/fileserver"
	"fmt"
	"net/http"
)

func GetIngresso(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("accessToken")
	if err != nil {
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		return
	}
	if cookie.Value == "-1" {
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		return
	}

	email := DB.GetEmail(cookie.Value)
	standing := DB.GetStanding(email)
	if !standing {
		Erros.HttpError(w, "dashboard/ingresso", fmt.Errorf("Usuário não possui email verificado"))
		return
	}

	var t = fileserver.Execute("template/ingresso.gohtml")
	t.Execute(w, nil)
}
