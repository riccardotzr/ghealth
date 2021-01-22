package ghealth

// HealthCheckDependency struct represents the dependency to check
type HealthCheckDependency struct {
	Name    string
	Checker HealthCheckInterface
}

// HealthCheckAggregator struct aggregate a list of checkers
type HealthCheckAggregator struct {
	checkers []HealthCheckDependency
}

// AddHealthCheck add a Checker to the Aggregator
func (aggregator *HealthCheckAggregator) AddHealthCheck(name string, checker HealthCheckInterface) {
	aggregator.checkers = append(aggregator.checkers, HealthCheckDependency{Name: name, Checker: checker})
}
