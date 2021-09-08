package db

import (
	"database/sql"
	"fmt"
	"time"
)

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

func NewUser(username, password, email, handle string) error {
	user := struct {
		Username string
		Password string
		Email    string
		Handle   string
	}{Username: username, Password: password, Email: email, Handle: handle}
	query := fmt.Sprintf("insert into users (username, email, password, handle) values ('%s', '%s', '%s', '%s');", user.Username, user.Email, user.Password, user.Handle)
	db := GetDB()
	_, err := db.Exec(query)
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}

func GetUserByEmail(email string) (UserWithoutNullString, error) {
	query := fmt.Sprintf("select * from users u where u.email='%s';", email)
	db := GetDB()
	var user User
	row := db.QueryRow(query)
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

func GetUserLikedTweets(userId, offset, limit int) ([]Tweet, error) {
	if offset < 0 {
		offset = 0
	}
	var tweets []Tweet
	db := GetDB()
	query := fmt.Sprintf("select tweet_id from likes l where l.who_liked=%d limit %d offset %d;", userId, limit, offset)
	rows, err := db.Query(query)
	if err != nil {
		return tweets, err
	}
	for rows.Next() {
		var tweetId int
		var tweet Tweet
		rows.Scan(&tweetId)
		query2 := fmt.Sprintf("select * from tweets t where t.id=%d", tweetId)
		rows2 := db.QueryRow(query2)
		rows2.Scan(&tweet.Id, &tweet.Content, &tweet.Author, &tweet.Date)
		tweets = append(tweets, tweet)
	}
	if rows.Err() != nil {
		return tweets, err
	}
	return tweets, nil
}

/*
func GetFollowedUsers(userId int) ([]PublicUser, error) {
	db := GetDB()
	// * JOIN
	query := fmt.Sprintf("%d", 1)
	rows, err := db.Query(query)
	if err != nil {
		fmt.Println(err)
		return []PublicUser{}, err
	}
	for rows.Next() {

	}

	return []PublicUser{}, nil
}
*/
