package main

import (
	"github.com/betelgeuse-7/twitt/api"
	chi "github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func routes() *chi.Mux {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Get("/api/v1/tweet/{id}", api.GetTweet)
	r.Post("/api/v1/tweet", api.NewTweet)

	return r
}
