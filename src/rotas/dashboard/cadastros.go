package dashboard

import (
  "fmt"
  "strconv"
  "net/http"
  DB "SCTI/database"
  HTMX "SCTI/htmx"
)

func PostValidateEmail(w http.ResponseWriter, r *http.Request) {
  email := r.FormValue("Email")

  err := DB.MarkAsPaid(email)
  if err != nil {
    HTMX.Failure(w, "Falha ao validar o ingresso: ", err)
    return
  }

  HTMX.Success(w, "Ingresso validado com sucesso")
}

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

  email := DB.GetEmail(cookie.Value)
  if !DB.GetStanding(email) {
    HTMX.Failure(w, "Falha ao se cadastrar: ", fmt.Errorf("Usuário não possui email verificado"))
    return
  }

  paid, err := DB.IsUserPaid(cookie.Value)
  if err != nil {
    HTMX.Failure(w, "Falha ao se cadastrar: ", err)
    return
  }

  if !paid {
    HTMX.Failure(w, "Falha ao se cadastrar: ", fmt.Errorf("Usuário não possui ingresso validado"))
    return
  }

  activityID, err := strconv.Atoi(r.FormValue("id"))
  if err != nil {
    http.Error(w, fmt.Sprintf("Invalid activity ID: %v", err), http.StatusBadRequest)
    return
  }

  _, err = DB.SignupUserForActivity(cookie.Value, activityID)
  if err != nil {
    HTMX.Failure(w, "Falha ao se cadastrar: ", err)
    return
  }
  w.Header().Set("HX-Refresh", "true")
  w.WriteHeader(http.StatusOK)
}


func PostDescadastros(w http.ResponseWriter, r *http.Request) {
  cookie, err := r.Cookie("accessToken")
  if err != nil {
    fmt.Println("Error Getting cookie:", err)
    http.Redirect(w, r, "/login", http.StatusSeeOther)
    return
  }

  if cookie.Value == "-1" {
    fmt.Println("Invalid accessToken")
    http.Redirect(w, r, "/login", http.StatusSeeOther)
  }

  activityID, err := strconv.Atoi(r.FormValue("id"))
  if err != nil {
    HTMX.Failure(w, "Falha ao sair: ", err)
    return
  }

  err = DB.UnregisterUserFromActivity(cookie.Value, activityID)
  if err != nil {
    HTMX.Failure(w, "Falha ao sair: ", err)
    return
  }

  fmt.Println(activityID)
  w.Header().Set("HX-Refresh", "true")
  w.WriteHeader(http.StatusOK)
}

func RemoveRegisteredActivities(allActivities, registeredActivities []DB.Activity) []DB.Activity {
  registeredMap := make(map[int]bool)
  for _, activity := range registeredActivities {
    registeredMap[activity.Activity_id] = true
  }
  filteredActivities := make([]DB.Activity, 0, len(allActivities))
  for _, activity := range allActivities {
    if !registeredMap[activity.Activity_id] && activity.Activity_type == "MC" {
      filteredActivities = append(filteredActivities, activity)
    }
  }
  return filteredActivities
}
