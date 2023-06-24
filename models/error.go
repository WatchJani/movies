package models

import (
	"fmt"
	"log"

	"github.com/valyala/fasthttp"
)

func ErrorHandler(err error) {
	if err != nil {
		log.Println(err)
	}
}

func ParserError(err error, ctx *fasthttp.RequestCtx) bool {
	if err != nil {
		fmt.Println("Greška prilikom parsiranja JSON-a:", err)
		ctx.SetStatusCode(fasthttp.StatusBadRequest)
		ctx.SetBodyString("Neispravan JSON format")
		return true
	}

	return false
}
