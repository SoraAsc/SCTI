package auth

import (
  "SCTI/fileserver"
  DB "SCTI/database"
  Erros "SCTI/erros"
  "fmt"
  "net/http"
  "net/url"
)

func GetVerify(w http.ResponseWriter, r *http.Request) {
  code := r.URL.Query().Get("code")
  encodedEmail := r.URL.Query().Get("email")

  if code == "" || encodedEmail == "" {
    Erros.HttpError(w, "auth/verify", fmt.Errorf("Code or Email parameter is missing"))
    return
  }

  email, err := url.QueryUnescape(encodedEmail)
  if err != nil {
    Erros.HttpError(w, "auth/verify", fmt.Errorf("Invalid Email format"))
    return
  }

  DB_Code, err := DB.GetCodeByEmail(email)
  if err != nil {
    Erros.HttpError(w, "auth/verify", fmt.Errorf("Database error"))
    return
  }

  if DB_Code == "" {
    Erros.HttpError(w, "auth/verify", fmt.Errorf("User didn't have a verification code"))
    return
  }

  if DB_Code != code {
    Erros.HttpError(w, "auth/verify", fmt.Errorf("Invalid verification code"))
    return
  }

  err = DB.SetStanding(DB.GetUUID(email), true)
  if err != nil {
    Erros.LogError("auth/login", fmt.Errorf("Error setting standing: %v", err))
  }

  var t = fileserver.Execute("template/verify.gohtml")
  t.Execute(w, nil)
}

func GetDelete(w http.ResponseWriter, r *http.Request) {
  code := r.URL.Query().Get("code")
  encodedEmail := r.URL.Query().Get("email")

  if code == "" || encodedEmail == "" {
    Erros.HttpError(w, "auth/verify", fmt.Errorf("Code or Email parameter is missing"))
    return
  }

  email, err := url.QueryUnescape(encodedEmail)
  if err != nil {
    Erros.HttpError(w, "auth/verify", fmt.Errorf("Invalid Email format"))
    return
  }

  DB_Code, err := DB.GetCodeByEmail(email)
  if err != nil {
    Erros.HttpError(w, "auth/verify", fmt.Errorf("Database error"))
    return
  }

  if DB_Code == "" {
    Erros.HttpError(w, "auth/verify", fmt.Errorf("User didn't have a verification code"))
    return
  }

  if DB_Code != code {
    Erros.HttpError(w, "auth/verify", fmt.Errorf("Invalid verification code"))
    return
  }

  uuid := DB.GetUUID(email)
  err = DB.DeleteUser(uuid)
  if err != nil {
    Erros.HttpError(w, "auth/verify", fmt.Errorf("Error deleting user: %v", err.Error()))
    return
  }

  var t = fileserver.Execute("template/delete.gohtml")
  t.Execute(w, nil)
}
