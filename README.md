<div align="center">

# Ghealth

[![Build Status](https://github.com/riccardotzr/ghealth/workflows/build/badge.svg)](https://github.com/riccardotzr/ghealth/actions)
[![Go Report Card](https://goreportcard.com/badge/github.com/riccardotzr/ghealth)](https://goreportcard.com/report/github.com/riccardotzr/ghealth)
[![Go Reference](https://pkg.go.dev/badge/github.com/riccardotzr/ghealth.svg)](https://pkg.go.dev/github.com/riccardotzr/ghealth)

</div>

GHealth is a minimal go library to check the health status of your application.

## Install

```ssh
go get -u github.com/riccardotzr/ghealth
```

## Usage

The library provides an interface: **HealthCheckInterface** that every health check must implement in order to check the health of the application.

```go

type HTTPHealthCheck struct {
	URL     string
	Timeout time.Duration
}

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

```

In this way it's possible to implement different health checks such as SQL, MongoDb, Kafka, RabbitMQ and so on. Do not implement this type of controls within the library is by design. Each application has different requirements and different dependencies. In this way it's possible to be flexible but open to any customization.

After that it's necessary to instantiate the handler who will take care of managing the various health checks, in any route you decide.

```go

handler := HealthCheckHandler{}
handler.AddHealthCheck("HTTP", &HTTPHealthCheck{URL: "https://www.google.com/", Timeout: 5 * time.Second})

```

The handler works as a checks aggregator. You can add as many checks as you need, as long as they implement the interface described above.

By default the handler returns an **HTTP Status Code 200** if all dependencies are healthy and an **HTTP Status Code 503** if at least one dependency is not healthy.
The json returned is the following:

```json
{
    "status": "Healthy",
    "dependencies": [
        {
            "name": "HTTP",
            "status": "Healthy"
        }
    ]
}
```


## License

This project is licensed under the Apache License 2.0 - see the [LICENSE.md](LICENSE.md)
file for details
