package auth

import (
  "SCTI/fileserver"
  DB "SCTI/database"
  "fmt"
  "net/http"
  "net/url"
)

func GetVerify(w http.ResponseWriter, r *http.Request) {
  code := r.URL.Query().Get("code")
  encodedEmail := r.URL.Query().Get("email")

  if code == "" || encodedEmail == "" {
    http.Error(w, "Code or email parameter is missing", http.StatusBadRequest)
    return
  }

  email, err := url.QueryUnescape(encodedEmail)
  if err != nil {
    http.Error(w, "Invalid email format", http.StatusBadRequest)
    return
  }

  DB_Code, err := DB.GetCodeByEmail(email)
  if err != nil {
    http.Error(w, "Database error", http.StatusInternalServerError)
    return
  }

  if DB_Code == "" {
    http.Error(w, "User didn't have a verification code", http.StatusUnauthorized)
    return
  }

  if DB_Code != code {
    http.Error(w, "Invalid verification code", http.StatusUnauthorized)
    return
  }

  err = DB.SetStanding(DB.GetUUID(email), true)
  if err != nil {
    fmt.Printf("Error setting standing: %v", err)
  }

  var t = fileserver.Execute("template/verify.gohtml")
  t.Execute(w, nil)
}

func GetDelete(w http.ResponseWriter, r *http.Request) {
  code := r.URL.Query().Get("code")
  encodedEmail := r.URL.Query().Get("email")

  if code == "" || encodedEmail == "" {
    http.Error(w, "Code or email parameter is missing", http.StatusBadRequest)
    return
  }

  email, err := url.QueryUnescape(encodedEmail)
  if err != nil {
    http.Error(w, "Invalid email format", http.StatusBadRequest)
    return
  }

  DB_Code, err := DB.GetCodeByEmail(email)
  if err != nil {
    http.Error(w, "Database error", http.StatusInternalServerError)
    return
  }

  if DB_Code == "" {
    http.Error(w, "User didn't have a verification code", http.StatusUnauthorized)
    return
  }

  if DB_Code != code {
    http.Error(w, "Invalid verification code", http.StatusUnauthorized)
    return
  }

  uuid := DB.GetUUID(email)
  err = DB.DeleteUser(uuid)
  if err != nil {
    http.Error(w, fmt.Sprintf("Error deleting user: %v", err.Error()), http.StatusUnauthorized)
    return
  }

  var t = fileserver.Execute("template/delete.gohtml")
  t.Execute(w, nil)
}
