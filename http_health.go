package ghealth

import (
	"net/http"
	"time"
)

// HTTPHealthCheck struct to configure HTTP Health Check
type HTTPHealthCheck struct {
	URL     string
	Timeout time.Duration
}

// CheckHealth is a implementation of HealthCheckInterface
func (healthCheck *HTTPHealthCheck) CheckHealth() HealthCheckResult {
	client := http.Client{
		Timeout: healthCheck.Timeout,
	}

	response, err := client.Head(healthCheck.URL)

	if response != nil {
		defer response.Body.Close()
	}

	if err != nil {
		return HealthCheckResult{status: Unhealthy}
	}

	if response.StatusCode != http.StatusOK {
		return HealthCheckResult{status: Unhealthy}
	}

	return HealthCheckResult{status: Healthy}
}

var _ HealthCheckInterface = &HTTPHealthCheck{}
