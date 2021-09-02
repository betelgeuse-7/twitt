package main

import (
	"github.com/betelgeuse-7/twitt/api"
	"github.com/betelgeuse-7/twitt/api/api_middleware"
	chi "github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func routes() *chi.Mux {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(api_middleware.ContentTypeJSON)
	r.Use(api_middleware.AuthorizationMiddleware)

	tweetsRoute := chi.NewRouter()
	tweetsRoute.Get("/{id}", api.GetTweet)
	tweetsRoute.Post("/", api.NewTweet)
	tweetsRoute.Post("/{id}/like", api.LikeTweet)
	tweetsRoute.Delete("/{id}", api.DeleteTweet)

	usersRoute := chi.NewRouter()
	usersRoute.Post("/new", api.Register)
	usersRoute.Post("/login", api.Login)
	// ?offset={offset}&limit={limit}
	usersRoute.Get("/{id}/liked", api.LikedTweets)

	r.Mount("/api/v1/tweet", tweetsRoute)
	r.Mount("/api/v1/users", usersRoute)
	return r
}
