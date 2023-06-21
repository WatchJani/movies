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

const (
	GET_ALL_ACTORS = "SELECT * FROM movies.actor;"
)

func GetAllActors() *[]Actor {
	var actors []Actor

	ErrorHandler(db.Select(&actors, GET_ALL_ACTORS))

	return &actors
}
