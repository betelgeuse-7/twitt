package db

import "database/sql"

var db *sql.DB

func GetDB() *sql.DB {
	return db
}
