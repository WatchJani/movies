package models

import (
	"github.com/WatchJani/movies.git/db"
)

type Genre struct {
	GenreID          uint   `db:"GenreID"`
	GenreName        string `db:"GenreName"`
	GenreDescription string `db:"GenreDescription"`
}

const (
	GET_ALL_GENRE = "SELECT * FROM heroku_d0920392b7eace1.genre;"
)

func GetAllGenre() *[]Genre {
	var genres []Genre

	ErrorHandler(db.Select(&genres, GET_ALL_GENRE))

	return &genres
}
