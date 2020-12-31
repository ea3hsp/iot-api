package service

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/httplog"
)

func InitService() {
	// Logger
	logger := httplog.NewLogger("iot-api-log", httplog.Options{
		JSON: true,
	})

	// Service
	r := chi.NewRouter()
	r.Use(httplog.RequestLogger(logger))

	r.Get("/domo", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("this is domo"))
	})

	r.Get("/test", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("this is test"))
	})

	http.ListenAndServe(":5000", r)
}
