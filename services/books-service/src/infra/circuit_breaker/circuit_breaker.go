package circuitbreaker

import (
	"time"

	"github.com/sony/gobreaker"
)

func NewCircuitBreakerInstance() *gobreaker.CircuitBreaker {
	st := gobreaker.Settings{
		Name:        "integrationCircuitBreaker",
		MaxRequests: 20,
		Interval:    2 * time.Second,
		Timeout:     40 * time.Second,
		ReadyToTrip: func(counts gobreaker.Counts) bool {
			return counts.ConsecutiveFailures > 3
		},
	}

	cb := gobreaker.NewCircuitBreaker(st)
	return cb
}

// MaxRequests:
// Description: The maximum number of requests that can be serviced
// before the Circuit Breaker breaks the circuit. If the number of requests exceeds this limit,
// the Circuit Breaker may temporarily disconnect the circuit.

// Interval:
// Description: The interval time between two measurements to check whether
// the circuit should be opened or closed again.
// This interval serves as a time window in which the failure is measured to determine
// whether the circuit should be opened or not.

// Timeout:
// Description: The maximum time allowed to wait for a
// a request before it is considered a failure.
// If the time spent to complete the request exceeds this timeout, the request is considered a failure.

// ReadyToTrip:
// Description: This is a function that determines whether or not a circuit should be opened.
// This function accepts the counts parameter which contains information about the number of consecutive failures,
// and then returns a boolean value. If the returned value is true, the circuit will be opened.
