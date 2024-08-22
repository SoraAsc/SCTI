package dashboard

import (
  "fmt"
  "strconv"
  "net/http"
  DB "SCTI/database"
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

  state, err := DB.SignupUserForActivity(cookie.Value, activityID)
  if err != nil {
    w.WriteHeader(http.StatusBadRequest)
  }

  fmt.Println(state)

  fmt.Println(activityID)
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
        if !registeredMap[activity.Activity_id] {
            filteredActivities = append(filteredActivities, activity)
        }
    }

    return filteredActivities
}
