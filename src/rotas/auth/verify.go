package auth

import (
  "SCTI/fileserver"
	DB "SCTI/database"
	"fmt"
	"net/http"
)

func GetVerify(w http.ResponseWriter, r *http.Request) {
  code := r.URL.Query().Get("code")
  uuid := r.URL.Query().Get("uuid")

  if code == "" || uuid == "" {
    http.Error(w, "Code or UUID parameter is missing", http.StatusBadRequest)
    return
  }

  DB_Code, err := DB.GetCode(uuid)
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

  err = DB.SetStanding(uuid, true)
  if err != nil {
    fmt.Printf("Error setting standing: %v", err)
  }

  var t = fileserver.Execute("template/verify.gohtml")
  t.Execute(w, nil)
}
