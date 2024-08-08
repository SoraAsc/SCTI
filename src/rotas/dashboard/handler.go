package dashboard

import (
  "net/http"
  "SCTI/fileserver"
)

type Handler struct{}

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

func RegisterRoutes(mux *http.ServeMux) {
  handler := &Handler{}
  mux.HandleFunc("GET /dashboard", handler.GetDashboard)
  mux.HandleFunc("GET /logoff", handler.GetLogoff)
}
func(h *Handler) GetLogoff(w http.ResponseWriter, r *http.Request) {
    http.SetCookie(w, &http.Cookie{
      Name:   "accessToken",
      Value:  "",
      MaxAge: -1,
      Secure: false,
      HttpOnly: true,
      Path: "/",
      SameSite: http.SameSiteLaxMode,
    })
    http.Redirect(w, r, "/login", http.StatusSeeOther)

    w.WriteHeader(http.StatusOK)
  }
