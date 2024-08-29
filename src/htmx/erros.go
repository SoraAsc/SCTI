package htmx

import (
  "net/http"
)

func Failure(w http.ResponseWriter, message string, err error) {
  w.Header().Set("Content-Type", "text/html")
  w.Write([]byte(`
  <div>
  ` + message + `
  ` + err.Error() + `
  </div>
  `))
}

func Success(w http.ResponseWriter, message string) {
  w.Header().Set("Content-Type", "text/html")
  w.Write([]byte(`
  <div>
  ` + message + `
  </div>
  `))
}
