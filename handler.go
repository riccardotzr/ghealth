package ghealth

import (
	"encoding/json"
	"net/http"
)

// HealtCheckResponseItem struct is the json response that indicates the health status of each dependencies
type HealtCheckResponseItem struct {
	Name   string `json:"name"`
	Status string `json:"status"`
}

// HealthCheckResponse struct is the json response returned by the health route
type HealthCheckResponse struct {
	Status       string                   `json:"status"`
	Dependencies []HealtCheckResponseItem `json:"dependencies"`
}

// HealthCheckHandler is a HTTP Server Handler Implementation
type HealthCheckHandler struct {
	HealthCheckAggregator
}

// NewHealthCheckHandler returns a new handler
func NewHealthCheckHandler() HealthCheckHandler {
	return HealthCheckHandler{}
}

// ServerHTTP returns a json encoded health response
func (handler HealthCheckHandler) ServeHTTP(response http.ResponseWriter, request *http.Request) {
	response.Header().Add("Content-Type", "application/json")

	var result HealthCheckResponse

	for _, check := range handler.HealthCheckAggregator.checkers {
		checkResult := check.Checker.CheckHealth()
		result.Dependencies = append(result.Dependencies, HealtCheckResponseItem{Name: check.Name, Status: checkResult.status})
	}

	var isUnhealthy = contains(result.Dependencies, Unhealthy)

	if isUnhealthy {
		result.Status = Unhealthy

		response.WriteHeader(http.StatusServiceUnavailable)
		json.NewEncoder(response).Encode(result)
	} else {
		result.Status = Healthy

		response.WriteHeader(http.StatusOK)
		json.NewEncoder(response).Encode(result)
	}
}
