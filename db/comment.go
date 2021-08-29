package db

import "time"

type Comment struct {
	Id      uint      `json:"comment_id"`
	Content string    `json:"content"`
	Tweet   Tweet     `json:"tweet"`
	Author  User      `json:"author"`
	Date    time.Time `json:"date"`
}
