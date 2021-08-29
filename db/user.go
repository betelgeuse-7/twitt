package db

import "time"

type User struct {
	Id           uint      `json:"user_id"`
	Username     string    `json:"username"`
	Password     string    // ! do not expose this
	Email        string    // ! do not expose this
	Handle       string    `json:"handle"`
	RegisterDate time.Time `json:"register_date"`
	Location     string    `json:"location"`
	Bio          string    `json:"bio"`
}
