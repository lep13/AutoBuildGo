package healthcheck

type HealthChecker interface {
	GetHealthStatus() Status
}

type RealHealthChecker struct{}

func (hc RealHealthChecker) GetHealthStatus() Status {
	// Implementation here
	return Status{Healthy: true, Message: "OK"}
}

// Status struct definition
type Status struct {
	Healthy bool
	Message string
}
