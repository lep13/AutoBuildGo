package middleware

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRecoveryFunc(t *testing.T) {
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		panic("test panic")
	})
	testHandler := RecoveryFunc(handler)

	req := httptest.NewRequest(http.MethodGet, "/", nil)
	res := httptest.NewRecorder()

	testHandler.ServeHTTP(res, req)

	assert.Equal(t, http.StatusOK, res.Code) // Assuming the handler recovers and returns StatusOK
}

func TestSetContentTypeJsonFunc(t *testing.T) {
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("test response"))
	})
	testHandler := SetContentTypeJsonFunc(handler)

	req := httptest.NewRequest(http.MethodGet, "/", nil)
	res := httptest.NewRecorder()

	testHandler.ServeHTTP(res, req)

	assert.Equal(t, "application/json", res.Header().Get("Content-Type"))
	assert.Equal(t, "test response", res.Body.String())
}
