package app

import (
	"context"
	"github.com/ChatService/internal/handler/http"
)

type StandaloneServer interface {
	Start(ctx context.Context)
}

type standaloneServer struct {
	httpServer http.HttpServer
}

func (s *standaloneServer) Start(ctx context.Context) {
	s.httpServer.Start(ctx)
}

func NewStandaloneServer(
	httpServer http.HttpServer,
) StandaloneServer {
	return &standaloneServer{
		httpServer: httpServer,
	}
}
