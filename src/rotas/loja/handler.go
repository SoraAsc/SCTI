package loja

import (
  "fmt"
  "net/http"
)

type Handler struct{}

func (h *Handler) GetLoja(w http.ResponseWriter, r *http.Request) {
  fmt.Fprintf(w, "Loja")
}

func RegisterRoutes(mux *http.ServeMux) {
  handler := &Handler{}
  mux.HandleFunc("GET /loja", handler.GetLoja)
}
