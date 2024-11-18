package route

import (
	"net/http"

	handlers "learn-microservices/src/interface/rest/handlers/books_grpc"

	"github.com/go-chi/chi/v5"
)

// HealthRouter a completely separate router for health check routes
func BookGRPCRouter(h handlers.BooksHandlerInterface) http.Handler {
	r := chi.NewRouter()

	r.Get("/", h.GetBySubject)

	return r
}
