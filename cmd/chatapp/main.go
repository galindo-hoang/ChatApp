package main

import (
	"context"
	"github.com/ChatService/internal/app"
)

func main() {
	server, err := app.InitializeStandaloneServer("")
	if err != nil {
		panic(err)
	}

	server.Start(context.Background())
}
