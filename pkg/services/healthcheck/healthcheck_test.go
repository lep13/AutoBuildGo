package healthcheck

import (
	"reflect"
	"testing"
)

func TestRealHealthChecker_GetHealthStatus(t *testing.T) {
	checker := RealHealthChecker{}
	expectedStatus := Status{Healthy: true, Message: "OK"}

	t.Run("Check Health Status", func(t *testing.T) {
		actualStatus := checker.GetHealthStatus()
		if !reflect.DeepEqual(actualStatus, expectedStatus) {
			t.Errorf("RealHealthChecker.GetHealthStatus() = %v, want %v", actualStatus, expectedStatus)
		}
	})
}
