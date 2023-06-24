package models

import (
	"github.com/WatchJani/movies.git/db"
)

type Movie struct {
	MovieID          uint   `db:"MovieID"`
	MovieTitle       string `db:"MovieTitle"`
	MovieDescription string `db:"MovieDescription"`
	MovieCover       string `db:"MovieCover"`
}

const (
	GET_ALL_MOVIES   = "SELECT * FROM heroku_d0920392b7eace1.movie;"
	INSERT_NEW_MOVIE = "INSERT INTO heroku_d0920392b7eace1.movie (MovieTitle, MovieDescription, MovieCover) VALUES (?, ?, ?);"
)

func GetAllMovies() *[]Movie {
	var movies []Movie
	ErrorHandler(db.Select(&movies, GET_ALL_MOVIES))

	return &movies
}

func InsertMovie(title, description, cover string) int64 {
	movie, err := db.Insert(INSERT_NEW_MOVIE, title, description, cover)

	ErrorHandler(err)

	movieID, err := movie.LastInsertId()

	ErrorHandler(err)

	return movieID
}
