package router

import (
	"github.com/Golang-Eskar/subscription-aggregator/internal/handlers"
	"github.com/Golang-Eskar/subscription-aggregator/internal/middleware"
	"github.com/go-chi/chi/v5"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func New() *chi.Mux {
	r := chi.NewRouter()
	r.Use(middleware.MetricsMiddleware)
	r.Handle("/metrics", promhttp.Handler())
	r.Post("/subscriptions", handlers.Create)
	r.Get("/subscriptions", handlers.GetAll)

	r.Get("/subscriptions/{id}", handlers.Get)
	r.Put("/subscriptions/{id}", handlers.Update)
	r.Delete("/subscriptions/{id}", handlers.Delete)

	r.Get("/subscriptions/filter", handlers.Filter)
	r.Get("/subscriptions/total", handlers.Sum)

	return r
}
