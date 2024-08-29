package dashboard

import (
	DB "SCTI/database"
	"html/template"
	"net/http"
)

type DashboardData struct {
  IsVerified bool
  IsAdmin bool
  Activities []DB.Activity
  RegisteredActivities []DB.Activity
}

func GetDashboard(w http.ResponseWriter, r *http.Request) {
  cookie, err := r.Cookie("accessToken")
  if err != nil {
    http.Redirect(w, r, "/login", http.StatusSeeOther)
    return
  }
  if cookie.Value == "-1" {
    http.Redirect(w, r, "/login", http.StatusSeeOther)
    return
  }

  all_activities, _ := DB.GetAllActivities()
  registered_activities, _ := DB.GetUserActivities(cookie.Value)
  available_activities := RemoveRegisteredActivities(all_activities, registered_activities)
  email := DB.GetEmail(cookie.Value)
  standing := DB.GetStanding(email)

  admin := false
  admcookie, err := r.Cookie("Admin")
  if err == nil && admcookie.Value == cookie.Value {
    admin = true
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

  data := DashboardData{
    IsVerified:          standing,
    IsAdmin:             admin,
    Activities:          available_activities,
    RegisteredActivities: registered_activities,
  }

  tmpl, err := template.ParseFiles("template/dashboard.gohtml")
  if err != nil {
    http.Error(w, err.Error(), http.StatusInternalServerError)
    return
  }
  err = tmpl.ExecuteTemplate(w, "dashboard", data)
  if err != nil {
    http.Error(w, err.Error(), http.StatusInternalServerError)
  }
}

func RegisterRoutes(mux *http.ServeMux) {
  mux.HandleFunc("GET /dashboard", GetDashboard)
  mux.HandleFunc("POST /cadastrar", PostCadastros)
  mux.HandleFunc("POST /descadastrar", PostDescadastros)
  mux.HandleFunc("POST /send-verification-email", VerifyEmail)
  mux.HandleFunc("POST /set-admin", SetAdmin)
  mux.HandleFunc("POST /add_activity", PostActivity)
}
