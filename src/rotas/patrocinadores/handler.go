package patrocinadores

import (
	"SCTI/fileserver"
	"net/http"
)

func GetPatrocinadores(w http.ResponseWriter, r *http.Request) {

	var t = fileserver.Execute("template/patrocinadores.gohtml")
	t.Execute(w, nil)
}

func RegisterRoutes(mux *http.ServeMux) {
	mux.HandleFunc("GET /patrocinadores", GetPatrocinadores)
}
