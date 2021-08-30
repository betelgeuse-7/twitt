package db

import (
	"database/sql"
	"fmt"
	"time"
)

type Tweet struct {
	Id      uint      `json:"tweet_id"`
	Content string    `json:"content"`
	Author  int       `json:"author_id"`
	Date    time.Time `json:"date"`
}

func GetTweetById(id int) Tweet {
	var tweet Tweet
	query := fmt.Sprintf("select * from tweets t where t.id = %d;", id)
	db := GetDB()
	row := db.QueryRow(query)
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
	query := fmt.Sprintf("insert into tweets (content, author) values ('%s', %d);", content, author)
	res, err := db.Exec(query)
	if err != nil {
		return 0, err
	}
	lastId, _ := res.LastInsertId()
	return int(lastId), nil
}
