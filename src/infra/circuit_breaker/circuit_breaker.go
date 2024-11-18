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
