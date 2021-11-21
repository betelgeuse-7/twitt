package db

import (
	"database/sql"
	"fmt"
	"time"
)

// i can just not put password and email fields.
// that's that easy
// |
// v

type User struct {
	Id           uint           `json:"user_id"`
	Username     string         `json:"username"`
	Password     string         // ! do not expose this
	Email        string         // ! do not expose this
	Handle       string         `json:"handle"`
	RegisterDate time.Time      `json:"register_date"`
	Location     sql.NullString `json:"location"`
	Bio          sql.NullString `json:"bio"`
}

type UserWithoutNullString struct {
	Id           uint      `json:"user_id"`
	Username     string    `json:"username"`
	Password     string    // ! do not expose this
	Email        string    // ! do not expose this
	Handle       string    `json:"handle"`
	RegisterDate time.Time `json:"register_date"`
	Location     string    `json:"location"`
	Bio          string    `json:"bio"`
}

type PublicUser struct {
	Id           uint      `json:"user_id"`
	Username     string    `json:"username"`
	Handle       string    `json:"handle"`
	RegisterDate time.Time `json:"register_date"`
	Location     string    `json:"location"`
	Bio          string    `json:"bio"`
}

// return a PublicUser with filled fields.
func (uwns UserWithoutNullString) publicUser() PublicUser {
	return PublicUser{
		Id:           uwns.Id,
		Username:     uwns.Username,
		Handle:       uwns.Handle,
		RegisterDate: uwns.RegisterDate,
		Location:     uwns.Location,
		Bio:          uwns.Bio,
	}
}

type followerStat struct {
	FollowerCount  int `json:"follower_count"`
	FollowingCount int `json:"following_count"`
}

func NewUser(username, password, email, handle string) error {
	user := struct {
		Username string
		Password string
		Email    string
		Handle   string
	}{Username: username, Password: password, Email: email, Handle: handle}
	db := GetDB()
	_, err := db.Exec("insert into users (username, email, password, handle) values ($1, $2, $3, $4)", user.Username, user.Email, user.Password, user.Handle)
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}

func GetUserBy(type_, value interface{}) (UserWithoutNullString, error) {
	db := GetDB()
	var row *sql.Row

	if type_ == "email" {
		row = db.QueryRow("select * from users u where u.email=$1", value)
	} else if type_ == "id" {
		row = db.QueryRow("select * from users u where u.id=$1", value)
	} else {
		panic("pass a valid identifier")
	}
	var user User
	err := row.Scan(&user.Id, &user.Username, &user.Password, &user.Email, &user.Handle, &user.RegisterDate, &user.Location, &user.Bio)
	if err != nil {
		fmt.Println(err)
		return UserWithoutNullString{}, err
	}
	userWithoutNullString := UserWithoutNullString{
		Id:           user.Id,
		Username:     user.Username,
		Password:     user.Password,
		Email:        user.Email,
		Handle:       user.Handle,
		RegisterDate: user.RegisterDate,
		Location:     user.Location.String,
		Bio:          user.Bio.String,
	}
	return userWithoutNullString, nil
}

// TODO refactor this
func GetUserLikedTweets(userId, offset, limit int) ([]Tweet, error) {
	if offset < 0 {
		offset = 0
	}
	var tweets []Tweet
	db := GetDB()
	rows, err := db.Query("select tweet_id from likes l where l.who_liked=$1 limit $2 offset $3", userId, limit, offset)
	if err != nil {
		return tweets, err
	}
	for rows.Next() {
		var tweetId int
		var tweet Tweet
		rows.Scan(&tweetId)

		rows := db.QueryRow("select * from tweets t where t.id=$1", tweetId)
		rows.Scan(&tweet.Id, &tweet.Content, &tweet.Author, &tweet.Date)

		tweets = append(tweets, tweet)
	}
	if rows.Err() != nil {
		return tweets, err
	}
	return tweets, nil
}

func GetFollowedUsers(userId int) ([]PublicUser, error) {
	var followedUsers []PublicUser
	db := GetDB()
	rows, err := db.Query("select id, username, email, password, handle, register_date, location, bio from users u inner join follows f on f.user_id=u.id and f.follower_id=$1", userId)
	if err != nil {
		fmt.Println(err)
		return []PublicUser{}, err
	}
	for rows.Next() {
		var user User
		err := rows.Scan(&user.Id, &user.Username, &user.Email, &user.Password, &user.Handle, &user.RegisterDate, &user.Location, &user.Bio)
		if err != nil {
			fmt.Println(err)
			return []PublicUser{}, err
		}

		u := UserWithoutNullString{
			Id:           user.Id,
			Username:     user.Username,
			Password:     user.Password,
			Email:        user.Email,
			Handle:       user.Handle,
			RegisterDate: user.RegisterDate,
			Location:     user.Location.String,
			Bio:          user.Bio.String,
		}.publicUser()
		followedUsers = append(followedUsers, u)
	}
	return followedUsers, nil
}

func GetUserFollowedBy(userId int) ([]PublicUser, error) {
	var followingUsers []PublicUser
	db := GetDB()
	rows, err := db.Query("select id, username, email, password, handle, register_date, location, bio from users u inner join follows f on f.follower_id=u.id and f.user_id=$1", userId)
	if err != nil {
		fmt.Println(err)
		return []PublicUser{}, err
	}
	for rows.Next() {
		var user User
		err := rows.Scan(&user.Id, &user.Username, &user.Email, &user.Password, &user.Handle, &user.RegisterDate, &user.Location, &user.Bio)
		if err != nil {
			fmt.Println(err)
			return []PublicUser{}, err
		}

		u := UserWithoutNullString{
			Id:           user.Id,
			Username:     user.Username,
			Password:     user.Password,
			Email:        user.Email,
			Handle:       user.Handle,
			RegisterDate: user.RegisterDate,
			Location:     user.Location.String,
			Bio:          user.Bio.String,
		}.publicUser()
		followingUsers = append(followingUsers, u)
	}
	return followingUsers, nil
}

func GetFollowCounts(userId int) (followerStat, error) {
	panic("(GetFollowCounts) not implemented")
}

/* // ! BAD
func GetFollowCounts(userId int) (followerStat, error) {
	db := GetDB()
	followingCountQuery := fmt.Sprintf("select count(id) from users u inner join follows f on f.user_id = u.id and f.follower_id = %d;", userId)
	followerCountQuery := fmt.Sprintf("select count(id) from users u inner join follows f on f.follower_id = u.id and f.user_id = %d;", userId)

	var followStats followerStat
	followerCountRow := db.QueryRow(followerCountQuery)
	followerCountRow.Scan(&followStats.FollowerCount)
	if err := followerCountRow.Err(); err != nil {
		fmt.Println(err)
		return followerStat{}, err
	}
	followingCountRow := db.QueryRow(followingCountQuery)
	followingCountRow.Scan(&followStats.FollowingCount)
	if err := followingCountRow.Err(); err != nil {
		fmt.Println(err)
		return followerStat{}, err
	}
	return followStats, nil
}
*/
