package http

import (
	"context"
	"encoding/json"
	"net/http"
	"time"

	"github.com/go-chi/chi"

	"github.com/ea3hsp/iot-api/internal/api"
	chiMiddleware "github.com/go-chi/chi/middleware"
	"github.com/go-chi/cors"
	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

// NewHTTPServer creates a new http transport server
func NewHTTPServer(ctx context.Context, endpoints api.Endpoints) http.Handler {
	options := []httptransport.ServerOption{
		httptransport.ServerErrorEncoder(encodeError),
		httptransport.ServerBefore(httptransport.PopulateRequestContext),
	}
	// creates new router
	r := chi.NewRouter()
	// avoid crash and recover
	r.Use(chiMiddleware.Recoverer)
	// set timeout
	r.Use(chiMiddleware.Timeout(60 * time.Second))
	// CORS
	r.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"iot.espin.ovh"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	}))
	// Routes
	r.Route("/domo", func(r chi.Router) {
		r.Post("/telemetry", httptransport.NewServer(
			endpoints.CreateTelemetry,
			api.DecodeTelemetry,
			encodeResponse,
			options...,
		).ServeHTTP)
	})
	// prometheus metrics route and handler
	r.Handle("/metrics", promhttp.Handler())
	return r
}

func encodeResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	if err, ok := response.(error); ok && err != nil {
		encodeError(ctx, err, w)
		return nil
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	return json.NewEncoder(w).Encode(response)
}

func encodeError(_ context.Context, err error, w http.ResponseWriter) {
	if err == nil {
		panic("encodeError with nil error")
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(codeFrom(err))
	json.NewEncoder(w).Encode(map[string]string{
		"error": err.Error(),
	})
}

func codeFrom(err error) int {
	switch err {
	default:
		return http.StatusInternalServerError
	}
}
