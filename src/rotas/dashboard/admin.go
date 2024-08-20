package dashboard

import (
  DB "SCTI/database"
  "fmt"
  "net/http"
)

func SetAdmin(w http.ResponseWriter, r *http.Request) {
	email := r.FormValue("Email")

	err := DB.SetAdmin(DB.GetUUID(email), true)
  if err != nil {
    fmt.Printf("Error setting admin status: %v", err)
  }

	fmt.Printf("Success")
}
