package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/lep13/AutoBuildGo/pkg/services/healthcheck"
)

type handlerService struct{}

// New Service
func NewHandlerService() *handlerService {
	return &handlerService{}
}

// Implementation goes here...
func (service *handlerService) GetHealthStatus(w http.ResponseWriter, r *http.Request) {

	status := healthcheck.GetHealthStatus()

	w.WriteHeader(200)
	err := json.NewEncoder(w).Encode(status)
	if err != nil {
		w.WriteHeader(500)
		fmt.Fprintf(w, `{"error":"Internal server error"}`)
		return
	}

}
