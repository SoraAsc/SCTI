package dashboard

import (
	DB "SCTI/database"
	HTMX "SCTI/htmx"
	"fmt"
	"net/http"
	"os"
	"strconv"
	"time"
)

func CheckAdmin(w http.ResponseWriter, r *http.Request) bool {
	admcookie, err := r.Cookie("Admin")
	if err != nil {
		HTMX.Failure(w, "Falha ao ler cookie de Admin: ", err)
		return false
	}

	logincookie, err := r.Cookie("accessToken")
	if err != nil {
		HTMX.Failure(w, "Falha ao ler cookie de Login: ", err)
		return false
	}

	if admcookie.Value != logincookie.Value {
		HTMX.Failure(w, "Admin inválido: ", fmt.Errorf("Cookie de Login e Admin diferem no usuário"))
		http.SetCookie(w, &http.Cookie{
			Name:     "Admin",
			Value:    "",
			MaxAge:   -1,
			Secure:   false,
			HttpOnly: true,
			Path:     "/",
			SameSite: http.SameSiteLaxMode,
		})
		return false
	}
	return true
}

func SetAdmin(w http.ResponseWriter, r *http.Request) {
	if !CheckAdmin(w, r) {
		HTMX.Failure(w, "Endpoint exclusivo de admins", fmt.Errorf("Acesso proibido a usuários não admin"))
		return
	}
	email := r.FormValue("Email")

	err := DB.SetAdmin(DB.GetUUID(email), true)
	if err != nil {
		HTMX.Failure(w, "Falha ao criar o Admin: ", err)
		return
	}

	HTMX.Success(w, "Admin criado com sucesso")
}

func RemoveAdmin(w http.ResponseWriter, r *http.Request) {
	if !CheckAdmin(w, r) {
		HTMX.Failure(w, "Endpoint exclusivo de admins", fmt.Errorf("Acesso proibido a usuários não admin"))
		return
	}

	email := r.FormValue("Email")
	err := DB.SetAdmin(DB.GetUUID(email), false)
	if err != nil {
		HTMX.Failure(w, "Falha ao remover o Admin: ", err)
		return
	}

	HTMX.Success(w, "Admin removido com sucesso")
}

func PostActivity(w http.ResponseWriter, r *http.Request) {
	if !CheckAdmin(w, r) {
		HTMX.Failure(w, "Endpoint exclusivo de admins", fmt.Errorf("Acesso proibido a usuários não admin"))
		return
	}

	eventStart := os.Getenv("SCTI_START_DATE")
	hourMin := r.FormValue("time") + ":00"
	day, _ := strconv.Atoi(r.FormValue("day"))

	eventStartDate, _ := time.Parse(time.DateOnly, eventStart)
	activityHour, _ := time.Parse(time.TimeOnly, hourMin)
	activityTime := eventStartDate.AddDate(0, 0, day-1)
	activityTime = activityTime.Add((time.Hour * time.Duration(activityHour.Hour())) + (time.Hour * 3))

	var a DB.Activity
	a.Spots, _ = strconv.Atoi(r.FormValue("spots"))
	a.Activity_type = r.FormValue("type")
	a.Room = r.FormValue("room")
	a.Speaker = r.FormValue("speaker")
	a.Topic = r.FormValue("topic")
	a.Description = r.FormValue("description")
	a.Time = hourMin
	a.Day = day
	a.Timestamp = activityTime.Unix()

	_, err := DB.CreateActivity(a)
	if err != nil {
		HTMX.Failure(w, "Falha ao criar atividade", err)
		return
	}
	HTMX.Success(w, "Atividade criada com sucesso")
}
