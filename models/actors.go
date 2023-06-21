package models

import "time"

type Actor struct {
	ActorID       uint
	ActorName     string
	ActorLastName string
	ActorImage    string
	ActorBirth    time.Time
}

func GetAllActor() *[]Actor {
	var actors []Actor

	return &actors
}
