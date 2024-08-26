package dashboard

import (
  "fmt"
  "strconv"
  "net/http"
  DB "SCTI/database"
)

func AtividadeCheia(w http.ResponseWriter, err error) {
  w.Header().Set("Content-Type", "text/html")
  w.Write([]byte(`
      <div class="failure">
        Falha ao se cadastrar: 
    ` + err.Error() + `
      </div>
  `))
}

func PaidError(w http.ResponseWriter, err error) {
  w.Header().Set("Content-Type", "text/html")
  w.Write([]byte(`
      <div class="failure">
        Falha ao validar: 
    ` + err.Error() + `
      </div>
  `))
}

func PostValidateEmail(w http.ResponseWriter, r *http.Request) {
  email := r.FormValue("Email")

  err := DB.MarkAsPaid(email)
  if err != nil {
    PaidError(w, err)
    return
  }

  w.Header().Set("Content-Type", "text/html")
  w.Write([]byte(`
    <div>
      Ingresso Validado com sucesso!
    </div>
  `))
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
    AtividadeCheia(w, fmt.Errorf("Usuário não possui email verificado"))
    return
  }

  paid, err := DB.IsUserPaid(cookie.Value)
  if err != nil {
    AtividadeCheia(w, err)
    return
  }

  if !paid {
    AtividadeCheia(w, fmt.Errorf("Usuário não possui ingresso validado"))
    return
  }

  activityID, err := strconv.Atoi(r.FormValue("id"))
  if err != nil {
    http.Error(w, fmt.Sprintf("Invalid activity ID: %v", err), http.StatusBadRequest)
    return
  }

  state, err := DB.SignupUserForActivity(cookie.Value, activityID)
  if err != nil {
    AtividadeCheia(w, err)
    return
  }

  fmt.Println(state)

  fmt.Println(activityID)
  w.Header().Set("HX-Refresh", "true")
  w.WriteHeader(http.StatusOK)
}

func PostDescadastros(w http.ResponseWriter, r *http.Request) {
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

  err = DB.UnregisterUserFromActivity(cookie.Value, activityID)
  if err != nil {
    w.WriteHeader(http.StatusBadRequest)
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
