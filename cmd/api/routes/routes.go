package routes

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/lep13/AutoBuildGo/cmd/api/routes/internal/handlers"
	"github.com/lep13/AutoBuildGo/cmd/api/routes/internal/middleware"
	"github.com/lep13/AutoBuildGo/pkg/services/healthcheck"
)

func NewRouter() *mux.Router {

	healthChecker := healthcheck.RealHealthChecker{}

	service := handlers.NewHandlerService(healthChecker)

	router := mux.NewRouter()

	// Endpoints
	serviceRouter := router.PathPrefix("/api/v1").Subrouter()
	serviceRouter.Methods(http.MethodGet).Path("/health").HandlerFunc(service.GetHealthStatus)

	// Middleware
	router.Use(middleware.RecoveryFunc)
	router.Use(middleware.SetContentTypeJsonFunc)

	return router
}
