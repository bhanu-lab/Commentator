package transport

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

type Handler struct {
	Router *mux.Router
}

// NewHandler -returns pointer to Handler
func NewHandler() *Handler {
	return &Handler{}
}

/*
SetupRoutes - prepares handling paths and its functions
*/
func (h *Handler) SetupRoutes() {
	fmt.Println("Setting up Routes")
	h.Router = mux.NewRouter()
	h.Router.HandleFunc("/health/check", HealthCheck)
}

/*
HealthCheck - deals with /health/check API
*/
func HealthCheck(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "I am alive!! \n")
}
