package auth

import (
  DB "SCTI/database"
  "SCTI/fileserver"
  "encoding/json"
  "net/http"
  "time"
  "log"
  "fmt"
)

func (h *Handler) GetLogin(w http.ResponseWriter, r *http.Request) {
  cookie, err := r.Cookie("accessToken")
  if err == nil && cookie != nil {
    http.Redirect(w, r, "/dashboard", http.StatusSeeOther)
    return
  }

  var t = fileserver.Execute("template/login.gohtml")
  t.Execute(w, nil)
}

func (h *Handler) PostLogin(w http.ResponseWriter, r *http.Request) {
  var user User
  if r.Header.Get("Content-type") == "application/json" {
    err := json.NewDecoder(r.Body).Decode(&user)
    if err != nil {
      log.Fatal(err)
    }
  } else {
    if err := r.ParseForm(); err != nil {
      fmt.Println("r.Form dentro if: ", r.Form)
      log.Fatal(err)
    }
    user.Email = r.FormValue("Email")
    user.Password = r.FormValue("Senha")
  }
  login, uuid := VerifyLogin(user, w)
  if login {
    cookie := http.Cookie{
      Name: "accessToken",
      Value: uuid,
      Expires: time.Now().Add(2 * 24 * time.Hour),
      Secure: false,
      HttpOnly: true,
      Path: "/",
      SameSite: http.SameSiteLaxMode,
    }
    http.SetCookie(w, &cookie)
    http.Redirect(w, r, "/dashboard", http.StatusSeeOther)
  }
}

func VerifyLogin(user User, w http.ResponseWriter)(login bool, uuid string) {
  found, err := DB.UserExists(user.Email)

  if err != nil {
    println("DB check: Query Failed")
    return false, ""
  }
  
  if !found {
    // println("Verify Login: User not found")
    return false, ""
  }

  uuid = fmt.Sprint(DB.GetUUID(user.Email))
  if CheckPasswordHash(user.Password, DB.GetHash(user.Email)) && uuid != "" {
    return true, uuid
  }
  return false, ""
}


