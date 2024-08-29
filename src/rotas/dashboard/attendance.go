package dashboard

import (
	DB "SCTI/database"
  HTMX "SCTI/htmx"
	"fmt"
	"html/template"
	"net/http"
	"net/url"
	"strconv"
)

type AttendanceData struct {
  Activities []DB.Activity
  Uuid string
}

func GetAttendance(w http.ResponseWriter, r *http.Request) {
  if !CheckAdmin(w, r) {
    http.Error(w, "Endpoint exclusiva de Admins", http.StatusUnauthorized)
    HTMX.Failure(w, "Acesso proibido", fmt.Errorf("Não foi encontrado ou não é valido o cookie de admin"))
    return
  }

  code := r.URL.Query().Get("code")
  encodedEmail := r.URL.Query().Get("email")

  fmt.Println(code, encodedEmail)

  if code == "" || encodedEmail == "" {
    http.Error(w, fmt.Sprintf("Código ou email do usuário ausentes!\nCódigo: %v\nEmail: %v", code, encodedEmail), http.StatusBadRequest)
  }

  email, err := url.QueryUnescape(encodedEmail)
  if err != nil {
    http.Error(w, fmt.Sprintf("Invalid email format: %v", err.Error()), http.StatusBadRequest)
    return
  }
  fmt.Println(email)

  uuid := DB.GetUUID(email)
  fmt.Println(uuid)
  userActivities, err := DB.GetUserActivities(uuid)
  if err != nil {
    http.Error(w, fmt.Sprintf("Não foi possivel recuperar os cadastros do usuário!\n%v", err.Error()), http.StatusInternalServerError)
  }
  attendedActivities, err := DB.GetUserAttendedActivities(uuid)
  if err != nil {
    http.Error(w, fmt.Sprintf("Não foi possivel recuperar as presenças do usuário!\n%v", err.Error()), http.StatusInternalServerError)
  }
  userActivities = RemoveAttendedActivities(userActivities, attendedActivities)

  fmt.Println("Got activities")

  data := AttendanceData{
    Activities: userActivities,
    Uuid: uuid,
  }

  tmpl, err := template.ParseFiles("template/attendance.gohtml")
  if err != nil {
    http.Error(w, err.Error(), http.StatusInternalServerError)
    return
  }
  tmpl.ExecuteTemplate(w, "attendance", data)
}

func PostAttendance(w http.ResponseWriter, r *http.Request) {
  if !CheckAdmin(w, r) {
    http.Error(w, "Endpoint exclusiva de Admins", http.StatusUnauthorized)
    HTMX.Failure(w, "Acesso proibido", fmt.Errorf("Não foi encontrado ou não é valido o cookie de admin"))
    return
  }

  uuid := r.FormValue("Uuid")
  atividade := r.FormValue("Atividade")

  id, err := strconv.Atoi(atividade)
  if err != nil {
    HTMX.Failure(w, "Falha ao resgatar atividade: ", err)
    return
  }

  err = DB.MarkUserAttendance(uuid, id)
  if err != nil {
    HTMX.Failure(w, "Falha ao marcar presença do usuário: ", err)
    return
  }

  HTMX.Success(w, "Presença do usuário marcada com sucesso")
}

func RemoveAttendedActivities(registeredActivities []DB.Activity, attendedActivities []DB.Activity) []DB.Activity {
    attendedMap := make(map[int]bool)
    for _, activity := range attendedActivities {
        attendedMap[activity.Activity_id] = true
    }

    var remainingActivities []DB.Activity
    for _, activity := range registeredActivities {
        if !attendedMap[activity.Activity_id] {
            remainingActivities = append(remainingActivities, activity)
        }
    }

    return remainingActivities
}
