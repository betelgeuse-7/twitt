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
	tweetsRoute.Get("/{id}/like", api.LikeTweet)
	tweetsRoute.Delete("/{id}", api.DeleteTweet)
	tweetsRoute.Get("/{id}/comments", api.GetCommentsForTweet)
	tweetsRoute.Post("/comments/new", api.NewCommentForTweet)
	tweetsRoute.Get("/", api.GetTweetsWithHashtag)

	usersRoute := chi.NewRouter()
	usersRoute.Post("/new", api.Register)
	usersRoute.Post("/login", api.Login)
	usersRoute.Get("/{id}/liked", api.LikedTweets)
	usersRoute.Get("/{id}/feed", api.UserFeed)
	usersRoute.Get("/{id}", api.UserProfile)
	usersRoute.Get("/follows", api.UserFollowing)
	usersRoute.Get("/followed_by", api.UserFollowedBy)
	usersRoute.Get("/{id}/liked", api.UserLikedTweets)
	usersRoute.Get("/follow/{id}", api.FollowUser)

	r.Mount("/api/v1/tweet", tweetsRoute)
	r.Mount("/api/v1/users", usersRoute)
	return r
}
