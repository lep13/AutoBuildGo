package routes

import (
	"net/http"

	"github.com/lep13/AutoBuildGo/cmd/api/routes/internal"

	"github.com/gorilla/mux"
)

func NewRouter() *mux.Router {

	service := handlers.NewHandlerService()

	router := mux.NewRouter()

	//endpoints
	serviceRouter := router.PathPrefix("/api/v1").Subrouter()
	serviceRouter.Methods(http.MethodGet).Path("/health").HandlerFunc(service.GetHealthStatus)

	//middleware
	router.Use(middleware.RecoveryFunc)
	router.Use(middleware.SetContentTypeJsonFunc)

	return router
}
