package patrocionadores

import (
	"SCTI/fileserver"
	"net/http"
)

func GetPatrocionadores(w http.ResponseWriter, r *http.Request) {

	var t = fileserver.Execute("template/patrocionadores.gohtml")
	t.Execute(w, nil)
}

func RegisterRoutes(mux *http.ServeMux) {
	mux.HandleFunc("GET /patrocionadores", GetPatrocionadores)
}
