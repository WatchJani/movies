package models

import (
	"github.com/WatchJani/movies.git/db"
)

type Actor struct {
	ActorID       uint   `db:"actorid"`
	ActorName     string `db:"actorname"`
	ActorLastName string `db:"actorlastname"`
	ActorImage    string `db:"actorimage"`
}

type ActorMovies struct {
	MovieName string `db:"moviename"`
	ActorName string `db:"actorname"`
}

const (
	GET_ALL_ACTORS       = "SELECT * FROM movies.actor;"
	GET_ALL_ACTOR_MOVIES = "SELECT a.actorname, m.moviename FROM movies.actor AS a JOIN movies.movieactor AS ma ON a.actorid = ma.actorid JOIN movies.movies AS m ON ma.movieid = m.movieid WHERE a.actorname = $1"
)

func GetAllActors() *[]Actor {
	var actors []Actor

	ErrorHandler(db.Select(&actors, GET_ALL_ACTORS))

	return &actors
}

func GetAllActorMovies() *[]ActorMovies {
	var actorMovies []ActorMovies

	ErrorHandler(db.Select(&actorMovies, GET_ALL_ACTOR_MOVIES, "Tom"))

	return &actorMovies
}
