package db

import "time"

type User struct {
	Id           uint      `json:"user_id"`
	Username     string    `json:"username"`
	Handle       string    `json:"handle"`
	RegisterDate time.Time `json:"register_date"`
	Location     string    `json:"location"`
	Bio          string    `json:"bio"`
}

func (u User) GetFollows() []User {
	// SQL
	return []User{}
}

func (u User) GetAllLikedPosts() []Tweet {
	// SQL
	return []Tweet{}
}
