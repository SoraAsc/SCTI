package lncc

import (
	"SCTI/fileserver"
	"net/http"
)

type Handler struct{}

func (h *Handler) GetLncc(w http.ResponseWriter, r *http.Request) {
    var t = fileserver.Execute("template/lncc.gohtml")
    t.Execute(w, nil)
}

func RegisterRoutes(mux *http.ServeMux) {
  handler := &Handler{}
  mux.HandleFunc("/lncc", handler.GetLncc)
}
