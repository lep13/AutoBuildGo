package handlers

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHandlers(t *testing.T) {

	req := httptest.NewRequest(http.MethodGet, "/", nil)
	res := httptest.NewRecorder()

	svc := NewHandlerService()
	svc.GetHealthStatus(res, req)

	assert.Equal(t, res.Code, 200)
}
