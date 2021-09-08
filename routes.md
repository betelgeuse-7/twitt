# Exposed enpoints
<hr>

## tweets
- GET /api/v1/tweet/:id
    &emsp;get tweet with id.
- GET /api/v1/tweet/?hashtag=:hashtag
   &emsp;get tweets with hashtag ':hashtag'
- POST /api/v1/tweet
    &emsp;new tweet
- POST /api/v1/tweet/:id/like
    &emsp;like a tweet
- DELETE /api/v1/tweet/:id
    &emsp;delete a tweet

## comments
- GET /api/v1/tweet/:id/comments
  &emsp;get comments for tweet
- POST /api/v1/tweet/comments/new
  &emsp;a new comment for tweet

## users
- GET /api/v1/users/:id/feed
    &emsp;feed with pagination (e.g. 20 limit)
- GET /api/v1/users/:id
    &emsp;user profile information
- GET /api/v1/users/:id/follows
    &emsp;get followed users
- GET /api/v1/users/:id/followed_by
    &emsp;get users who follow user :id
- GET /api/v1/users/:id/liked
    &emsp;get liked tweets
- POST /api/v1/users/follow/:id
    &emsp;follow user
- POST /api/v1/users/new
    &emsp;new user
- POST /api/v1/users/login
    &emsp;log in user
