package dashboard

import (
  DB "SCTI/database"
  "fmt"
  "net/http"
  "html/template"
)

type DashboardData struct {
  HTMLContent template.HTML
  IsVerified bool
  IsAdmin bool
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

  html, _ := ActivitiesList()

  admin := DB.GetAdmin(cookie.Value)
  email := DB.GetEmail(cookie.Value)
  standing := DB.GetStanding(email)
  htmlContent := template.HTML(html)

  data := DashboardData{
    IsVerified: standing,
    HTMLContent: htmlContent,
    IsAdmin: admin,
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
  mux.HandleFunc("POST /send-verification-email", VerifyEmail)
  mux.HandleFunc("POST /add_activity", PostActivity)
}
