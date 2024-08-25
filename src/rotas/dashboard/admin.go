package dashboard

import (
	DB "SCTI/database"
	"net/http"
	"strconv"
	"time"
)

func SetAdmin(w http.ResponseWriter, r *http.Request) {
  email := r.FormValue("Email")

  err := DB.SetAdmin(DB.GetUUID(email), true)
  if err != nil {
    AdmError(w, err)
    return
  }

  w.Header().Set("Content-Type", "text/html")
  w.Write([]byte(`
    <div>
      Admin criado com sucesso
    </div>
  `))
}

func ActiError(w http.ResponseWriter, err error) {
  w.Header().Set("Content-Type", "text/html")
  w.Write([]byte(`
    <div>
      Falha ao criar atividade:
      ` + err.Error() + `
    </div>
  `))
}

func AdmError(w http.ResponseWriter, err error) {
  w.Header().Set("Content-Type", "text/html")
  w.Write([]byte(`
    <div>
      Falha ao criar adm: 
      ` + err.Error() + `
    </div>
  `))
}

func PostActivity(w http.ResponseWriter, r* http.Request) {
  var a DB.Activity
  a.Spots, _ = strconv.Atoi(r.FormValue("spots"))
  a.Activity_type = r.FormValue("type")
  a.Room = r.FormValue("room")
  a.Speaker = r.FormValue("speaker")
  a.Topic = r.FormValue("topic")
  a.Description = r.FormValue("description")
  a.Time = r.FormValue("time")
  a.Day, _ = strconv.Atoi(r.FormValue("day"))

  _, err := DB.CreateActivity(a)
  if err != nil {
    ActiError(w, err)
    return
  }

  w.Header().Set("Content-Type", "text/html")
  w.Write([]byte(`
    <div>
      Atividade criada com sucesso.
    </div>
  `))
}

func VerifyAdmin(w http.ResponseWriter, r *http.Request) bool {
  admcookie, _ := r.Cookie("Admin")
  if admcookie != nil && admcookie.Value == "1" {
    return true
  }

  logincookie, err := r.Cookie("accessToken")
  if err != nil {
    // fmt.Println("Error Getting login cookie:", err)
    return false
  }

  isAdmin := DB.GetAdmin(logincookie.Value)
  if isAdmin {
    admcookie := http.Cookie{
      Name:     "Admin",
      Value:    "1",
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
