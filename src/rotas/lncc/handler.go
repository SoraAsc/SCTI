package lncc

import (
	"SCTI/fileserver"
	"net/http"
)

func GetLncc(w http.ResponseWriter, r *http.Request) {
	var t = fileserver.Execute("template/lncc.gohtml")
	t.Execute(w, nil)
}

func RegisterRoutes(mux *http.ServeMux) {
	mux.HandleFunc("/lncc", GetLncc)
}
