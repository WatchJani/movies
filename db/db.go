package db

import (
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

var db *sqlx.DB

func Init() error {
	var err error
	db, err = sqlx.Open("postgres", "user=janko dbname=movies password=JankoKondic72621@ sslmode=disable")
	if err != nil {
		return err
	}

	return nil
}

func Select(dest interface{}, query string, args ...interface{}) error {
	return db.Select(dest, query, args...)
}
