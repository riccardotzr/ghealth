package ghealth

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"gotest.tools/assert"
)

func Test_Handler_Server(t *testing.T) {

	t.Run("HealthCheck Up", func(t *testing.T) {
		expectedResult := HealthCheckResponse{
			Status: "Healthy",
			Dependencies: []HealtCheckResponseItem{
				{
					Name:   "HTTP",
					Status: "Healthy",
				},
			},
		}

		request, _ := http.NewRequest("GET", "/-/health", nil)
		response := httptest.NewRecorder()

		handler := HealthCheckHandler{}
		handler.AddHealthCheck("HTTP", &HTTPHealthCheck{URL: "https://www.google.com/", Timeout: 5 * time.Second})

		handler.ServeHTTP(response, request)

		responseBody, _ := ioutil.ReadAll(response.Body)

		var result HealthCheckResponse
		json.Unmarshal(responseBody, &result)

		assert.Equal(t, response.Code, http.StatusOK)
		assert.DeepEqual(t, result, expectedResult)
	})

	t.Run("HealthCheck Down", func(t *testing.T) {
		expectedResult := HealthCheckResponse{
			Status: "Unhealthy",
			Dependencies: []HealtCheckResponseItem{
				{
					Name:   "HTTP",
					Status: "Unhealthy",
				},
			},
		}

		request, _ := http.NewRequest("GET", "/-/health", nil)
		response := httptest.NewRecorder()

		handler := HealthCheckHandler{}
		handler.AddHealthCheck("HTTP", &HTTPHealthCheck{URL: "https://fakego.com", Timeout: 5 * time.Second})

		handler.ServeHTTP(response, request)

		responseBody, _ := ioutil.ReadAll(response.Body)

		var result HealthCheckResponse
		json.Unmarshal(responseBody, &result)

		assert.Equal(t, response.Code, http.StatusServiceUnavailable)
		assert.DeepEqual(t, result, expectedResult)
	})
}
