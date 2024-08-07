package dashboard

import (
  "fmt"
  "net/http"
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
  fmt.Fprintf(w, "User ID: %v", auth.Value)
}

func RegisterRoutes(mux *http.ServeMux) {
  handler := &Handler{}
  mux.HandleFunc("GET /dashboard", handler.GetDashboard)
}
