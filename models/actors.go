package models

import (
	"time"
)

type Actor struct {
	ActorID       uint
	ActorName     string
	ActorLastName string
	ActorImage    string
	ActorBirth    time.Time
}

// const (
// 	GET_ALL = "SELECT * FROM movies.movies;"
// )

// func GetAll() *[]Actor {
// 	var actors []Actor

// 	err := db.Select(&actors, GET_ALL)

// 	if err != nil {
// 		log.Println(err)
// 	}

// 	return &actors
// }
