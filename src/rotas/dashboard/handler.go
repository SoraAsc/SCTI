package dashboard

import (
	DB "SCTI/database"
	"fmt"
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


  email := DB.GetEmail(cookie.Value)
  standing := DB.GetStanding(email)

  var admin bool

  admcookie, err := r.Cookie("Admin")
  if err != nil && err.Error() != "http: named cookie not present" {
      http.Error(w, fmt.Sprintf("Invalid admin cookie: %v", err), http.StatusBadRequest)
      admin = false
  }

  logincookie, err := r.Cookie("acessToken")
  if err != nil {
    w.Header().Set("Content-Type", "text/html")
    w.Write([]byte(`
    <div class="failure">
    Falha ao ler cookie de login:
    ` + err.Error() + `
    </div>
    `))
    return
  }

  if admcookie.Value == logincookie.Value {
    admin = true
  } else {
    http.SetCookie(w, &http.Cookie{
      Name:   "Admin",
      Value:  "",
      MaxAge: -1,
      Secure: false,
      HttpOnly: true,
      Path: "/",
      SameSite: http.SameSiteLaxMode,
    })
  admin = false
  }

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

func RegisterRoutes(mux *http.ServeMux) {
  mux.HandleFunc("GET /dashboard", GetDashboard)
  mux.HandleFunc("POST /cadastrar", PostCadastros)
  mux.HandleFunc("POST /descadastrar", PostDescadastros)
  mux.HandleFunc("POST /send-verification-email", VerifyEmail)
  mux.HandleFunc("POST /set-admin", SetAdmin)
  mux.HandleFunc("POST /add_activity", PostActivity)
}
