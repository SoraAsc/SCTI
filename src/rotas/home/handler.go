package home

import (
    "SCTI/fileserver"
    "net/http"
)

type Handler struct{}

func (h *Handler) GetHome(w http.ResponseWriter, r *http.Request) {
    var t = fileserver.Execute("template/index.gohtml")
    t.Execute(w, nil)
}

func RegisterRoutes(mux *http.ServeMux) {
    handler := &Handler{}
    mux.HandleFunc("/", handler.GetHome)
}
