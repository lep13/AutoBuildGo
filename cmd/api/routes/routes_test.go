package routes

import (
    "net/http"
    "net/http/httptest"
    "testing"

    "github.com/stretchr/testify/assert"
    "github.com/gorilla/mux"
    "github.com/lep13/AutoBuildGo/cmd/api/routes/internal/middleware"
)

func TestNewRouter(t *testing.T) {
    router := mux.NewRouter()

    // Ensure the RecoveryFunc is the first middleware to catch any panics
    router.Use(middleware.RecoveryFunc)

    // Then apply other middleware
    router.Use(middleware.SetContentTypeJsonFunc)

    // Set up a route that triggers a panic
    router.Handle("/panic", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        panic("Test panic")
    }))

    // Testing the RecoveryFunc middleware
    t.Run("Test_Recovery_Middleware", func(t *testing.T) {
        req, _ := http.NewRequest("GET", "/panic", nil)
        resp := httptest.NewRecorder()
        router.ServeHTTP(resp, req)

        // The response should be 500 Internal Server Error
        assert.Equal(t, http.StatusInternalServerError, resp.Code, "Expected HTTP status 500 after panic")
        assert.Equal(t, "application/json", resp.Header().Get("Content-Type"), "Expected Content-Type to be application/json")
    })

    // Add the health check route for testing
    router.HandleFunc("/api/v1/health", func(w http.ResponseWriter, r *http.Request) {
        w.WriteHeader(http.StatusOK)
        w.Header().Set("Content-Type", "application/json")
        w.Write([]byte(`{"status":"ok"}`))
    })

    t.Run("Health Check Route", func(t *testing.T) {
        req, _ := http.NewRequest(http.MethodGet, "/api/v1/health", nil)
        resp := httptest.NewRecorder()
        router.ServeHTTP(resp, req)

        // Check the response of the health check
        assert.Equal(t, http.StatusOK, resp.Code, "Expected HTTP status 200 for health check")
        assert.Equal(t, "application/json", resp.Header().Get("Content-Type"), "Expected Content-Type to be application/json")
    })
}