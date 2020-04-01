package main

import (
	"fmt"
	"github.com/savsgio/atreugo/v11"
)

func main() {
	config := atreugo.Config{
		Addr: "127.0.0.1:8000",
	}

	server := atreugo.New(config)
	server.Path("GET", "/channels/{id}/programm", ProgrammAction)

	err := server.ListenAndServe()

	if err != nil {
		panic(err)
	}
}

func ProgrammAction(ctx *atreugo.RequestCtx) error {
	id := ctx.UserValue("id").(string)

	return ctx.TextResponse(fmt.Sprintf("Hello World by %s", id))
}
