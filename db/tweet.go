package db

import (
	"database/sql"
	"fmt"
	"time"
)

// ! https://golang.org/doc/tutorial/database-access#:~:text=By%20separating%20the%20SQL%20statement%20from%20parameter%20values%20(rather%20than%20concatenating%20them%20with%2C%20say%2C%20fmt.Sprintf)%2C%20you%20enable%20the%20database/sql%20package%20to%20send%20the%20values%20separate%20from%20the%20SQL%20text%2C%20removing%20any%20SQL%20injection%20risk.

type Tweet struct {
	Id      uint      `json:"tweet_id"`
	Content string    `json:"content"`
	Author  int       `json:"author_id"`
	Date    time.Time `json:"date"`
}

func GetTweetById(id int) Tweet {
	var tweet Tweet
	db := GetDB()
	row := db.QueryRow("select * from tweets t where t.id = $1", id)
	if row.Err() != nil {
		return tweet
	}
	err := row.Scan(&tweet.Id, &tweet.Content, &tweet.Author, &tweet.Date)
	if err != sql.ErrNoRows {
		return tweet
	}
	return tweet
}

func NewTweet(content string, author int) (int, error) {
	db := GetDB()
	res, err := db.Exec("insert into tweets (content, author) values ($1, $2)", content, author)
	if err != nil {
		return 0, err
	}
	lastId, _ := res.LastInsertId()
	return int(lastId), nil
}

func LikeTweet(tweetId, userId int) error {
	db := GetDB()
	_, err := db.Exec("insert into likes (tweet_id, who_liked) values ($1, $2)", tweetId, userId)
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}

func DeleteTweet(id int) error {
	db := GetDB()
	_, err := db.Exec("delete from tweets t where t.id=$1", id)
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}
