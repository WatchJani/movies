package models

import (
	"github.com/WatchJani/movies.git/db"
)

type Movie struct {
	MovieID          uint   `db:"movieid"`
	MovieName        string `db:"moviename"`
	MovieDescription string `db:"moviedescription"`
	MovieCover       string `db:"moviecover"`
}

const (
	GET_ALL_MOVIES = "SELECT * FROM movies.movies;"
)

func GetAllMovies() *[]Movie {
	var movies []Movie
	ErrorHandler(db.Select(&movies, GET_ALL_MOVIES))

	return &movies
}
