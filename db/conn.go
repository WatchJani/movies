package db

import (
	"github.com/jmoiron/sqlx"
)

var db *sqlx.DB

func Init() error {
	var err error
	db, err = sqlx.Open("postgres", "user=janko dbname=to-do password=JankoKondic72621@ sslmode=disable")
	if err != nil {
		return err
	}

	return nil
}
