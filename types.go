package ghealth

const (
	// Healthy Status
	Healthy = "Healthy"

	// Unhealthy Status
	Unhealthy = "Unhealthy"
)

// HealthCheckResult struct used to indicate the result of the health check
type HealthCheckResult struct {
	status string
}

// HealthCheckInterface used to check the health status of the dependency
type HealthCheckInterface interface {
	CheckHealth() HealthCheckResult
}
