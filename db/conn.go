package db

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

var database *sql.DB

type Postgres struct {
	Host     string
	Port     string
	User     string
	DbName   string
	Password string
	SslMode  bool
}

func (p Postgres) makeConnStr() string {
	sslmode := "disable"
	if p.SslMode {
		sslmode = "enable"
	}
	return fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=%s password=%s", p.Host, p.Port, p.User, p.DbName, sslmode, p.Password)
}

func (p Postgres) Connect() (*sql.DB, error) {
	db, err := sql.Open("postgres", p.makeConnStr())
	if err != nil {
		return nil, err
	}
	database = db
	return db, nil
}

func GetDB() *sql.DB {
	return database
}
