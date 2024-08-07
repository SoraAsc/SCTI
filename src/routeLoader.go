package main

import (
  "net/http"

  "SCTI/fileserver"
  "SCTI/rotas/auth"
  "SCTI/rotas/home"
  "SCTI/rotas/loja"
  "SCTI/rotas/lncc"
  "SCTI/rotas/dashboard"
)

func LoadRoutes(mux *http.ServeMux) {
  mux.Handle("/static/", http.StripPrefix("/static/", fileserver.FS))

  auth.RegisterRoutes(mux)
  dashboard.RegisterRoutes(mux)
  home.RegisterRoutes(mux)
  lncc.RegisterRoutes(mux)
  loja.RegisterRoutes(mux)
}
