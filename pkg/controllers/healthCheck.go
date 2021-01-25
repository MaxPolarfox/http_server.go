package controllers

import (
	"fmt"
	"net/http"
)

// HealthcheckController is a controller that handles /readiness and /liveness calls
type HealthcheckController struct {
}

// Liveness responds to /liveness
func (h *HealthcheckController) Liveness(rw http.ResponseWriter, req *http.Request) {
	rw.WriteHeader(http.StatusOK)
	fmt.Fprintf(rw, "%s", http.StatusText(http.StatusOK))
}

// Readiness responds to /readiness
func (h *HealthcheckController) Readiness(rw http.ResponseWriter, req *http.Request) {
	rw.WriteHeader(http.StatusOK)
	fmt.Fprintf(rw, "%s", http.StatusText(http.StatusOK))
}
