package dashboard

import (
  "strconv"
  "net/http"
  DB "SCTI/database"
)

func ActiError(w http.ResponseWriter, err error) {
  w.Header().Set("Content-Type", "text/html")
  w.Write([]byte(`
      <div>
          Falha ao criar atividade.
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
  }
}
