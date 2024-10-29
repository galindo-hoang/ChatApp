package main

import (
	"context"
	"github.com/ChatService/internal/wiring"
)

func main() {
	server, f, err := wiring.InitializeStandaloneServer("")
	if err != nil {
		f()
		panic(err)
	}
	defer f()

	server.Start(context.Background())
}
