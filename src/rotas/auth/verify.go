package auth

import (
  "SCTI/fileserver"
  "net/http"
)

func (h *Handler) GetVerify(w http.ResponseWriter, r *http.Request) {
  var t = fileserver.Execute("template/verify.gohtml")
  t.Execute(w, nil)
}

func (h *Handler) PostVerify(w http.ResponseWriter, r *http.Request) {
}
