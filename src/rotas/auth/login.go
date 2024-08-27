package auth

import (
  DB "SCTI/database"
  "SCTI/fileserver"
  "encoding/json"
  "net/http"
  "time"
  "fmt"
)

func LoginFailed(w http.ResponseWriter, err error) {
  w.Header().Set("Content-Type", "text/html")
  w.Write([]byte(`
      <div class="failure">
        Falha no Login:
    ` + err.Error() + `
      </div>
  `))
}

func GetLogin(w http.ResponseWriter, r *http.Request) {
  var t = fileserver.Execute("template/login.gohtml")
  t.Execute(w, nil)
}

func PostLogin(w http.ResponseWriter, r *http.Request) {
  var user User
  if r.Header.Get("Content-type") == "application/json" {
    err := json.NewDecoder(r.Body).Decode(&user)
    if err != nil {
      LoginFailed(w, err)
      return
    }
  } else {
    if err := r.ParseForm(); err != nil {
      LoginFailed(w, err)
      return
    }
    user.Email = r.FormValue("Email")
    user.Password = r.FormValue("Senha")
  }
  login, uuid, err := VerifyLogin(user, w)

  if err != nil {
    LoginFailed(w, err)
    return
  }

  if !login {
    LoginFailed(w, fmt.Errorf("Somehow login failed"))
    return
  }

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
  VerifyAdmin(w, r, uuid)

  if r.Header.Get("HX-Request") == "true" {
    w.Header().Set("HX-Redirect", "/dashboard")
    w.WriteHeader(http.StatusOK)
    return
  }

  http.Redirect(w, r, "/dashboard", http.StatusSeeOther)
}

func VerifyLogin(user User, w http.ResponseWriter)(login bool, uuid string, err error) {
  found, err := DB.UserExists(user.Email)

  if err != nil {
    return false, "", fmt.Errorf("DB check query Failed")
  }
 
  if !found {
    return false, "", fmt.Errorf("User not found")
  }

  uuid = fmt.Sprint(DB.GetUUID(user.Email))
  if CheckPasswordHash(user.Password, DB.GetHash(user.Email)) && uuid != "" {
    return true, uuid, nil
  }
  return false, "", fmt.Errorf("Senha inválida")
}

func GetLogoff(w http.ResponseWriter, r *http.Request) {
  http.SetCookie(w, &http.Cookie{
    Name:   "accessToken",
    Value:  "",
    MaxAge: -1,
    Secure: false,
    HttpOnly: true,
    Path: "/",
    SameSite: http.SameSiteLaxMode,
  })
  http.SetCookie(w, &http.Cookie{
    Name:   "Admin",
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

func VerifyAdmin(w http.ResponseWriter, r *http.Request, uuid string) bool {
  admcookie, err := r.Cookie("Admin")
  if err != nil && err.Error() != "http: named cookie not present" {
      http.Error(w, fmt.Sprintf("Invalid admin cookie: %v", err), http.StatusBadRequest)
      return false
  }

  if admcookie.Value == uuid {
    return true
  }

  isAdmin := DB.GetAdmin(uuid)
  fmt.Println("login.go valor do isAdmin:", isAdmin)
  if isAdmin {
    admcookie := http.Cookie{
      Name:     "Admin",
      Value:    uuid,
      Expires:  time.Now().Add(2 * 24 * time.Hour),
      Secure:   false,
      HttpOnly: true,
      Path:     "/",
      SameSite: http.SameSiteLaxMode,
    }
    http.SetCookie(w, &admcookie)
    return true
  }
  return false
}

