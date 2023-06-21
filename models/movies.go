package models

import (
	"log"

	"github.com/WatchJani/movies.git/db"
)

type Movie struct {
	MovieID          int    `db:"movieid"`
	MovieName        string `db:"moviename"`
	MovieDescription string `db:"moviedescription"`
	MovieCover       string `db:"moviecover"`
}

const (
	GET_ALL_MOVIES = "SELECT * FROM movies.movies;"
)

func ErrorHandler(err error) {
	if err != nil {
		log.Println(err)
	}
}

func GetAllMovies() *[]Movie {
	var movies []Movie
	ErrorHandler(db.Select(&movies, GET_ALL_MOVIES))

	return &movies
}
