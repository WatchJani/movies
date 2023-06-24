package handler

import (
	"encoding/json"
	"sync"

	"github.com/WatchJani/movies.git/models"
	"github.com/valyala/fasthttp"
)

func HandleCORS(h fasthttp.RequestHandler) fasthttp.RequestHandler {
	return fasthttp.RequestHandler(func(ctx *fasthttp.RequestCtx) {
		ctx.Response.Header.Set("Access-Control-Allow-Origin", "*")
		ctx.Response.Header.Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE")
		ctx.Response.Header.Set("Access-Control-Allow-Headers", "Content-Type")

		if ctx.IsOptions() {
			ctx.SetStatusCode(fasthttp.StatusOK)
			return
		}

		h(ctx)
	})
}

func GetMovies(ctx *fasthttp.RequestCtx) {
	jsonResponse, err := json.Marshal(models.GetAllMovies())
	if err != nil {
		ctx.SetStatusCode(fasthttp.StatusInternalServerError)
		return
	}

	ctx.SetContentType("application/json")
	ctx.SetBody(jsonResponse)
}

func GetActors(ctx *fasthttp.RequestCtx) {
	jsonResponse, err := json.Marshal(models.GetAllActors())
	if err != nil {
		ctx.SetStatusCode(fasthttp.StatusInternalServerError)
		return
	}

	ctx.SetContentType("application/json")
	ctx.SetBody(jsonResponse)
}

func GetGenre(ctx *fasthttp.RequestCtx) {
	jsonResponse, err := json.Marshal(models.GetAllGenre())
	if err != nil {
		ctx.SetStatusCode(fasthttp.StatusInternalServerError)
		return
	}

	ctx.SetContentType("application/json")
	ctx.SetBody(jsonResponse)
}

func GetActorMovies(ctx *fasthttp.RequestCtx) {
	actorID := ctx.UserValue("id").(string)

	jsonResponse, err := json.Marshal(models.GetAllActorMovies(actorID))
	if err != nil {
		ctx.SetStatusCode(fasthttp.StatusInternalServerError)
		return
	}

	ctx.SetContentType("application/json")
	ctx.SetBody(jsonResponse)
}

func GetHome(ctx *fasthttp.RequestCtx) {
	ctx.SetContentType("application/json")
	ctx.SetBody([]byte("Welcome to our server!"))
}

func InsertNewMovie(ctx *fasthttp.RequestCtx) {
	requestBody := ctx.PostBody()

	request := struct {
		MovieTitle       string
		MovieDescription string
		MovieCover       string
		MovieActors      []int
		MovieGenres      []int
	}{}

	err := json.Unmarshal(requestBody, &request)
	if models.ParserError(err, ctx) {
		return
	}

	movieID := models.InsertMovie(request.MovieTitle, request.MovieDescription, request.MovieCover)

	var wg sync.WaitGroup

	wg.Add(2)

	go models.InsertActorsInMovie(movieID, request.MovieActors, &wg)
	go models.InsertGenreInMovie(movieID, request.MovieGenres, &wg)

	wg.Wait()

	ctx.SetStatusCode(fasthttp.StatusCreated)
	ctx.SetBodyString(string(requestBody))
}

func InsertNewActor(ctx *fasthttp.RequestCtx) {
	requestBody := ctx.PostBody()

	newActor := struct {
		ActorName     string
		ActorLastName string
		ActorImage    string
	}{}

	err := json.Unmarshal(requestBody, &newActor)
	if models.ParserError(err, ctx) {
		return
	}

	models.ErrorHandler(models.InsertActor(newActor.ActorName, newActor.ActorLastName, newActor.ActorImage))

	ctx.SetStatusCode(fasthttp.StatusCreated)
	ctx.SetBodyString(string(requestBody))
}

func DeleteNewActor(ctx *fasthttp.RequestCtx) {
	actorID := ctx.UserValue("id").(string)

	models.ErrorHandler(models.DeleteActor(actorID))
	ctx.SetStatusCode(fasthttp.StatusNoContent) // TAKO KAZE NEW DA JE OVO STATUS ZA BRISANJE
}

func UpdateNewActor(ctx *fasthttp.RequestCtx) {
	requestBody := ctx.PostBody()
	actorID := ctx.UserValue("id").(string)

	actorNewData := struct {
		ActorName     string
		ActorLastName string
		ActorImage    string
	}{}

	err := json.Unmarshal(requestBody, &actorNewData)
	if models.ParserError(err, ctx) {
		return
	}

	models.ErrorHandler(models.UpdateActor(actorNewData.ActorName, actorNewData.ActorLastName, actorNewData.ActorImage, actorID))
	ctx.SetStatusCode(fasthttp.StatusOK)
}
