package main

import (
	"net/http"

	"SCTI/fileserver"
	"SCTI/rotas/auth"
	"SCTI/rotas/dashboard"
	"SCTI/rotas/home"
	"SCTI/rotas/lncc"
	"SCTI/rotas/loja"
	"SCTI/rotas/patrocionadores"
)

func LoadRoutes(mux *http.ServeMux) {
	mux.Handle("/static/", http.StripPrefix("/static/", fileserver.FS))

	auth.RegisterRoutes(mux)
	dashboard.RegisterRoutes(mux)
	home.RegisterRoutes(mux)
	lncc.RegisterRoutes(mux)
	loja.RegisterRoutes(mux)
	patrocionadores.RegisterRoutes(mux)
}
