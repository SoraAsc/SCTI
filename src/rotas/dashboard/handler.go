package dashboard

import (
  "SCTI/fileserver"
  "encoding/json"
  "fmt"
  "log"
  "net/http"
)

type Handler struct{}

type Courses struct {
  Seg string `json:"seg"`
  Ter string `json:"ter"`
  Qua string `json:"qua"`
  Qui string `json:"qui"`
  Sex string `json:"sex"`
}

func (h *Handler) GetDashboard(w http.ResponseWriter, r *http.Request) {
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
  var t = fileserver.Execute("template/dashboard.gohtml")
  t.Execute(w, nil)
  // fmt.Fprintf(w, "User ID: %v", auth.Value)
}

func (h *Handler) PostDashboard(w http.ResponseWriter, r *http.Request) {
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
  var courses Courses
  if r.Header.Get("Content-Type") == "application/json" {
    err := json.NewDecoder(r.Body).Decode(&courses)
    if err != nil {
      log.Printf("Error decoding JSON: %v", err)
      http.Error(w, "Invalid JSON", http.StatusBadRequest)
      return
    }
  } else {
    // Lida com application/x-www-form-urlencoded
    if err := r.ParseForm(); err != nil {
      log.Printf("Error parsing form: %v", err)
      http.Error(w, "Invalid form data", http.StatusBadRequest)
      return
    }
    courses.Seg = r.FormValue("seg")
    courses.Ter = r.FormValue("ter")
    courses.Qua = r.FormValue("qua")
    courses.Qui = r.FormValue("qui")
    courses.Sex = r.FormValue("sex")
    fmt.Fprintf(w, "%s\n%s\n%s\n%s\n%s", courses.Seg, courses.Ter, courses.Qua, courses.Qui, courses.Sex)
  }
}

func RegisterRoutes(mux *http.ServeMux) {
  handler := &Handler{}
  mux.HandleFunc("GET /dashboard", handler.GetDashboard)
  mux.HandleFunc("POST /dashboard", handler.PostDashboard)
}
