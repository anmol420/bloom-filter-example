package main

import (
	"log"
	"net/http"
	"time"

	"github.com/anmol420/bloom-filter-example/internal/store"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-playground/validator/v10"
)

var Validate *validator.Validate

func init() {
	Validate = validator.New(validator.WithRequiredStructEnabled())
}

type application struct {
	config config
	store  store.Storage
}

type config struct {
	port string
	db   dbConfig
}

type dbConfig struct {
	dbURI  string
	dbName string
}

func (app *application) mount() http.Handler {
	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.Timeout(60 * time.Second))

	r.Route("/api/v1", func(r chi.Router) {
		r.Get("/health", app.healthCheckHandler)
		r.Route("/users", func(r chi.Router) {
			r.Post("/create", app.createUserHandler)
			r.Get("/search/{username}", app.searchUserHandler)
		})
	})

	return r
}

func (app *application) run(mux http.Handler) error {
	srv := &http.Server{
		Addr:         app.config.port,
		Handler:      mux,
		WriteTimeout: time.Second * 30,
		ReadTimeout:  time.Second * 10,
		IdleTimeout:  time.Minute,
	}
	log.Printf("Server Started On %s", app.config.port)
	return srv.ListenAndServe()
}
