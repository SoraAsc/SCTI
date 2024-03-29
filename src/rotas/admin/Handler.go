package admin

import (
	"fmt"
	"net/http"
)

type Handler struct{}

func (h *Handler) GetDashboard(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Admin dashboard")
}
