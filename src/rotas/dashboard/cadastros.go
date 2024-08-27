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

  admcookie, err := r.Cookie("Admin")
  if err != nil {
    w.Header().Set("Content-Type", "text/html")
    w.Write([]byte(`
    <div class="failure">
    Falha ao ler cookie de admin:
    ` + err.Error() + `
    </div>
    `))
    return
  } else {
    admin, err := strconv.ParseBool(admcookie.Value)
    if err != nil {
      w.Header().Set("Content-Type", "text/html")
      w.Write([]byte(`
      <div class="failure">
      Falha ao converter cookie de admin:
      ` + err.Error() + `
      </div>
      `))
      return
    }
    if admin {
      email := DB.GetEmail(cookie.Value)
      if !DB.GetStanding(email) {
        AtividadeCheia(w, fmt.Errorf("Usuário não possui email verificado"))
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
      w.WriteHeader(http.StatusOK)}
    }
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

    admcookie, err := r.Cookie("Admin")
    if err != nil {
      w.Header().Set("Content-Type", "text/html")
      w.Write([]byte(`
      <div class="failure">
      Falha ao ler cookie de admin:
      ` + err.Error() + `
      </div>
      `))
      return
    } else {
      admin, err := strconv.ParseBool(admcookie.Value)
      if err != nil {
        w.Header().Set("Content-Type", "text/html")
        w.Write([]byte(`
        <div class="failure">
        Falha ao converter cookie de admin:
        ` + err.Error() + `
        </div>
        `))
        return
      }
      if admin {
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
    }
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
