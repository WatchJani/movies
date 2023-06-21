package main

import (
	"fmt"
	"log"

	"github.com/WatchJani/movies.git/db"
	"github.com/WatchJani/movies.git/models"
	_ "github.com/lib/pq"
)

func main() {
	err := db.Init()

	if err != nil {
		log.Println("Database is not connected!")
	}

	fmt.Println(models.GetAllMovies())
}
