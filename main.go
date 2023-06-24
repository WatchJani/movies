package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/WatchJani/movies.git/db"
	"github.com/WatchJani/movies.git/handler"
	"github.com/fasthttp/router"
	"github.com/valyala/fasthttp"
)

func makeRequest() {
	resp, err := http.Get("https://golang-movies-2c7ce157bd9b.herokuapp.com/")
	if err != nil {
		fmt.Println("Greška prilikom izvršavanja zahtjeva:", err)
		return
	}
	defer resp.Body.Close()
}

func main() {
	PORT := os.Getenv("PORT")

	ticker := time.NewTicker(10 * time.Minute)

	go func() {
		for range ticker.C {
			makeRequest()
		}
	}()

	if db.Init() != nil {
		log.Println("Database is not connected!")
	}

	app := router.New()

	app.GET("/", handler.GetHome)
	app.GET("/movies", handler.GetMovies)
	app.GET("/actors", handler.GetActors)
	app.GET("/genres", handler.GetGenre)
	app.GET("/movies/actor/{id}", handler.GetActorMovies)
	app.POST("/movie", handler.InsertNewMovie)
	app.POST("/actor", handler.InsertNewActor)
	app.DELETE("/actor/{id}", handler.DeleteNewActor)
	app.PUT("/actor/{id}", handler.UpdateNewActor)

	log.Fatal(fasthttp.ListenAndServe(":"+PORT, handler.HandleCORS(app.Handler)))
}
