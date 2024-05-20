package healthcheck

import (
	"reflect"
	"testing"
)

func TestRealHealthChecker_GetHealthStatus(t *testing.T) {
	tests := []struct {
		name string
		hc   RealHealthChecker
		want Status
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.hc.GetHealthStatus(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("RealHealthChecker.GetHealthStatus() = %v, want %v", got, tt.want)
			}
		})
	}
}
