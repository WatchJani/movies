package models

import (
	"github.com/WatchJani/movies.git/db"
)

type Actor struct {
	ActorID       uint   `db:"ActorID"`
	ActorName     string `db:"ActorName"`
	ActorLastName string `db:"ActorLastName"`
	ActorImage    string `db:"ActorImage"`
}

const (
	GET_ALL_ACTORS   = "SELECT * FROM heroku_d0920392b7eace1.actor;"
	INSERT_NEW_ACTOR = "INSERT INTO `heroku_d0920392b7eace1`.`actor` (`ActorName`, `ActorLastName`, `ActorImage`) VALUES (?,?,?)"
	DELETE_ACTOR     = "DELETE FROM heroku_d0920392b7eace1.actor WHERE actorID = ?;"
	PUT_ACTOR        = "UPDATE heroku_d0920392b7eace1.actor SET ActorName = ?, ActorLastName = ?, ActorImage = ? WHERE ActorID = ?;"
)

func GetAllActors() *[]Actor {
	var actors []Actor

	ErrorHandler(db.Select(&actors, GET_ALL_ACTORS))

	return &actors
}

func InsertActor(name, lastName, image string) error {
	_, err := db.Insert(INSERT_NEW_ACTOR, name, lastName, image)

	return err
}

func DeleteActor(actorID string) error {
	_, err := db.Delete(DELETE_ACTOR, actorID)

	return err
}

func UpdateActor(actorName, actorLastName, actorImage, actorID string) error {
	_, err := db.Update(PUT_ACTOR, actorName, actorLastName, actorImage, actorID)

	return err
}
