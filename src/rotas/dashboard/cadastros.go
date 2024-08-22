package dashboard

import (
  "fmt"
  "strconv"
  "net/http"
)

func PostCadastros(w http.ResponseWriter, r *http.Request) {
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
  activityID, err := strconv.Atoi(r.FormValue("id"))
  if err != nil {
    http.Error(w, fmt.Sprintf("Invalid activity ID: %v", err), http.StatusBadRequest)
    return
  }

  fmt.Println(activityID)
  w.WriteHeader(http.StatusOK)
}

