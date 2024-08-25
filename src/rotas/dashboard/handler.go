package dashboard

import (
  DB "SCTI/database"
  "fmt"
  "net/http"
  "html/template"
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
    // fmt.Println("Error Getting cookie:", err)
    http.Redirect(w, r, "/login", http.StatusSeeOther)
    return
  }

  if cookie.Value == "-1" {
    // fmt.Println("Invalid accessToken")
    http.Redirect(w, r, "/login", http.StatusSeeOther)
  }

  all_activities, _ := DB.GetAllActivities()
  registered_activities, _ := DB.GetUserActivities(cookie.Value)
  available_activities := RemoveRegisteredActivities(all_activities, registered_activities)

  admin := VerifyAdmin(w, r)
  email := DB.GetEmail(cookie.Value)
  standing := DB.GetStanding(email)

  data := DashboardData{
    IsVerified: standing,
    IsAdmin: admin,
    Activities: available_activities,
    RegisteredActivities: registered_activities,
  }

  tmpl, err := template.ParseFiles("template/dashboard.gohtml")
  if err != nil {
    http.Error(w, err.Error(), http.StatusInternalServerError)
    return
  }
  tmpl.ExecuteTemplate(w, "dashboard", data)
}

func PostDashboard(w http.ResponseWriter, r *http.Request) {
  auth, err := r.Cookie("accessToken")
  if err != nil {
    // fmt.Println("Error Getting cookie:", err)
    http.Redirect(w, r, "/login", http.StatusSeeOther)
    return
  }

  if auth.Value == "-1" {
    // fmt.Println("Invalid accessToken")
    http.Redirect(w, r, "/login", http.StatusSeeOther)
  }
  fmt.Fprintf(w, "POST /dashboard")
}

func RegisterRoutes(mux *http.ServeMux) {
  mux.HandleFunc("GET /dashboard", GetDashboard)
  mux.HandleFunc("POST /dashboard", PostDashboard)
  mux.HandleFunc("POST /cadastrar", PostCadastros)
  mux.HandleFunc("POST /descadastrar", PostDescadastros)
  mux.HandleFunc("POST /send-verification-email", VerifyEmail)
  mux.HandleFunc("POST /set-admin", SetAdmin)
  mux.HandleFunc("POST /add_activity", PostActivity)
}
