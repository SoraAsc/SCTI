package htmx

import (
	"net/http"
)

func Failure(w http.ResponseWriter, message string, err error) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.Write([]byte(`
  <div class="modal">
  <div class="error-messages">
  ` + message + `
  ` + err.Error() + `
  </div>
   <span class="close" onclick=" const modals = document.querySelectorAll('.modal');
        if (modals.length > 0) {
            modals[modals.length - 1].style.display = 'none';
        }">sair</span>
  </div>

  `))
}

func Success(w http.ResponseWriter, message string) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.Write([]byte(`
  <div class="modal">
  <div class = "success-messages">
  ` + message + `
  </div>
  <span class="close" onclick=" const modals = document.querySelectorAll('.modal');
        if (modals.length > 0) {
            modals[modals.length - 1].style.display = 'none';
        }">sair</span>
  </div>
  `))
}
