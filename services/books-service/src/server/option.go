package server

import (
	"learn-microservices-server/src/infra/config"
	
)

// WithConfig is function
func WithConfig(config *config.Config) ServerGrpcOption {
	return func(r *Server) {
		r.config = config
	}
}

