package handlers

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/lep13/AutoBuildGo/pkg/services/healthcheck"
	"github.com/stretchr/testify/assert"
)

type mockHealthChecker struct {
	shouldFail bool
}

func (m *mockHealthChecker) GetHealthStatus() healthcheck.Status {
	if m.shouldFail {
		panic("simulated failure")
	}
	return healthcheck.Status{Healthy: true, Message: "System is healthy"}
}

func TestGetHealthStatus(t *testing.T) {
	cases := []struct {
		name          string
		shouldFail    bool
		expectedCode  int
		expectedBody  string
		expectedPanic bool
	}{
		{"success", false, http.StatusOK, "System is healthy", false},
		{"failure", true, http.StatusInternalServerError, "Internal server error", true},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			checker := &mockHealthChecker{shouldFail: c.shouldFail}
			service := NewHandlerService(checker)
			recorder := httptest.NewRecorder()
			request := httptest.NewRequest(http.MethodGet, "/health", nil)

			if c.expectedPanic {
				assert.PanicsWithValue(t, "simulated failure", func() {
					service.GetHealthStatus(recorder, request)
				}, "Expected panic did not occur")
			} else {
				service.GetHealthStatus(recorder, request)
				assert.Equal(t, c.expectedCode, recorder.Code)
				assert.Contains(t, recorder.Body.String(), c.expectedBody)
			}
		})
	}
}
