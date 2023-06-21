package models

import "github.com/WatchJani/movies.git/db"

type Genre struct {
	GenreID          uint   `db:"genreid"`
	GenreName        string `db:"genrename"`
	GenreDescription string `db:"genredescription"`
}

const (
	GET_ALL_GENRE = "SELECT * FROM movies.genre;"
)

func GetAllGenre() *[]Genre {
	var genres []Genre

	ErrorHandler(db.Select(&genres, GET_ALL_GENRE))

	return &genres
}
