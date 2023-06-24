package models

import (
	"sync"

	"github.com/WatchJani/movies.git/db"
)

const (
	INSERT_GENRE_IN_MOVIES = "INSERT INTO heroku_d0920392b7eace1.moviegenre (MovieID, GenreID)"
)

func InsertGenreInMovie(movieID int64, genres []int, wg *sync.WaitGroup) {
	insertQuery := GenerateCustomInput(INSERT_GENRE_IN_MOVIES, genres, movieID)

	_, err := db.Insert(insertQuery)

	ErrorHandler(err)

	wg.Done()
}
