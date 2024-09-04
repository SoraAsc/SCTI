package patrocionadores

import (
	"SCTI/rotas/notfound"
	"html/template"
	"net/http"
)

func GetPatrocionadores(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/patrocionadores" {
		notfound.NotFound(w, r)
		return
	}

	tmpl, err := template.ParseFiles("template/patrocionadores.gohtml")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	// Corrigindo o uso da função ExecuteTemplate
	err = tmpl.ExecuteTemplate(w, "patrocionadores", nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func RegisterRoutes(mux *http.ServeMux) {
	mux.HandleFunc("/patrocionadores", GetPatrocionadores)
}
