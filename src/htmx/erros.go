package htmx

import (
	"net/http"
)

func Failure(w http.ResponseWriter, message string, err error) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.Write([]byte(`


</style>
  <div class="error-messages" >
  <button  onclick="this.parentElement.style.display = 'none';">x</button>
  ` + message + `
  ` + err.Error() + `
  
  </div>

  `))
}

func Success(w http.ResponseWriter, message string) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.Write([]byte(`

  <div class = "success-messages">
   <button onclick="this.parentElement.style.display = 'none';">x</button>
  ` + message + `

  </div>
  `))
}
