package ingresso

import (
	"SCTI/fileserver"
	"net/http"
)

func GetIngresso(w http.ResponseWriter, r *http.Request) {

	var t = fileserver.Execute("template/ingresso.gohtml")
	t.Execute(w, nil)
}

func RegisterRoutes(mux *http.ServeMux) {
	mux.HandleFunc("GET /ingresso", GetIngresso)
}
