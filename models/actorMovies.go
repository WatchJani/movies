package models

import (
	"sync"

	"github.com/WatchJani/movies.git/db"
)

type ActorMovies struct {
	MovieName string `db:"MovieTitle"`
	ActorName string `db:"ActorName"`
}

const (
	GET_ALL_ACTOR_MOVIES    = "SELECT a.ActorName, m.MovieTitle FROM heroku_d0920392b7eace1.actor AS a JOIN heroku_d0920392b7eace1.movieactor AS ma ON a.ActorID = ma.ActorID JOIN heroku_d0920392b7eace1.movie AS m ON ma.MovieID = m.MovieID WHERE a.ActorID = ?;"
	INSERT_ACTORS_IN_MOVIES = "INSERT INTO heroku_d0920392b7eace1.movieactor (MovieID, ActorID)"
)

func GetAllActorMovies(id string) *[]ActorMovies {
	var actorMovies []ActorMovies

	ErrorHandler(db.Select(&actorMovies, GET_ALL_ACTOR_MOVIES, id))

	return &actorMovies
}

func InsertActorsInMovie(movieID int64, actors []int, wg *sync.WaitGroup) {
	insertQuery := GenerateCustomInput(INSERT_ACTORS_IN_MOVIES, actors, movieID)

	_, err := db.Insert(insertQuery)

	ErrorHandler(err)

	wg.Done()
}
