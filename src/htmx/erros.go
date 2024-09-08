package htmx

import (
	"net/http"
)

func Failure(w http.ResponseWriter, message string, err error) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.Write([]byte(`
  <div class="error-messages">
  ` + message + `
  ` + err.Error() + `
  </div>
  `))
}

func Success(w http.ResponseWriter, message string) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.Write([]byte(`
  <div class = "success-messages">
  ` + message + `
  </div>
  `))
}
