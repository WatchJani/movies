package main

import (
	"fmt"
	"log"
	"time"

	"github.com/WatchJani/movies.git/db"
	"github.com/WatchJani/movies.git/models"
	_ "github.com/lib/pq"
)

func main() {
	start := time.Now()

	err := db.Init()

	if err != nil {
		log.Println("Database is not connected!")
	}

	fmt.Println(models.GetAllMovies())
	fmt.Println(models.GetAllActors())
	fmt.Println(models.GetAllGenre())
	fmt.Println(models.GetAllActorMovies())

	fmt.Println(time.Since(start))
}
