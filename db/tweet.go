package db

import "time"

type Tweet struct {
	Id      uint      `json:"tweet_id"`
	Content string    `json:"content"`
	Author  User      `json:"author"`
	Date    time.Time `json:"date"`
}
