package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHomeRoutePing(t *testing.T) {
	// Cria um novo servidor HTTP de teste
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Aqui você colocaria o handler real da sua página inicial
		w.WriteHeader(http.StatusOK)
	}))
	defer ts.Close()

	// Faz uma requisição GET para o servidor de teste
	resp, err := http.Get(ts.URL)
	if err != nil {
		t.Fatalf("Falha ao fazer requisição: %v", err)
	}
	defer resp.Body.Close()

	// Verifica se o status code é 200 OK
	if resp.StatusCode != http.StatusOK {
		t.Errorf("Status code inesperado: recebido %v, esperado %v",
			resp.StatusCode, http.StatusOK)
	}
}
