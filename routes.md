## tweet
GET /api/v1/tweet/:id
    get tweet with id.
POST /api/v1/tweet
    new tweet
POST /api/v1/tweet/:id/like
    like a tweet
DELETE /api/v1/tweet/:id
    delete a tweet

## comments
GET /api/v1/tweet/:id/comments
    get comments for tweet
POST /api/v1/tweet/:id/comments
    a new comment for tweet
    
## users
GET /api/v1/users/:id/feed
    feed with pagination (e.g. 20 limit)
GET /api/v1/users/:id
    user profile information
GET /api/v1/users/:id/follows
    get followed users
GET /api/v1/users/:id/followed_by
    get users who follow user :id
GET /api/v1/users/:id/liked
    get liked tweets
POST /api/v1/users/follow/:id
    follow user
POST /api/v1/users/new
    new user
POST /api/v1/users/login
    log in user
