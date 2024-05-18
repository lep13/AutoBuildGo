package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/lep13/AutoBuildGo/pkg/services/healthcheck"
)

type HandlerService struct {
	HealthChecker healthcheck.HealthChecker
}

func NewHandlerService(hc healthcheck.HealthChecker) *HandlerService {
	return &HandlerService{HealthChecker: hc}
}

func (service *HandlerService) GetHealthStatus(w http.ResponseWriter, r *http.Request) {
	status := service.HealthChecker.GetHealthStatus()

	w.WriteHeader(http.StatusOK)
	err := json.NewEncoder(w).Encode(status)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintln(w, `{"error":"Internal server error"}`)
	}
}
