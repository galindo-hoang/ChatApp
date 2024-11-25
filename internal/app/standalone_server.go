package app

import (
	"context"
	"fmt"
	"github.com/ChatService/internal/configs"
	pb "github.com/ChatService/internal/generated/chat_app/v1"
	grpcHandler "github.com/ChatService/internal/handler/grpc"
	httpHandler "github.com/ChatService/internal/handler/http"
	"google.golang.org/grpc"
	"net"
	"net/http"
)

type StandaloneServer interface {
	Start(ctx context.Context)
}

type grpcServer struct {
	grpcConfig configs.GRPC
}

func (r *grpcServer) Start(ctx context.Context) {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%v", r.grpcConfig.Port))
	if err != nil {
		fmt.Printf("failed to listen: %v\n", err)
		return
	}
	s := grpc.NewServer()
	pb.RegisterRelationshipServiceServer(s, &grpcHandler.RelationshipHandler{})
	fmt.Printf("server listening at %v\n", lis.Addr())
	if err := s.Serve(lis); err != nil {
		fmt.Printf("failed to serve: %v\n", err)
	}
}

type httpServer struct {
	httpConfigs configs.Http
}

func (h *httpServer) Start(ctx context.Context) {

	http.HandleFunc("/v1/sessions", httpHandler.Middleware(httpHandler.CreateSession))
	http.HandleFunc("/v1/accounts", httpHandler.Middleware(httpHandler.CreateAccount))
	http.HandleFunc("/v1/sessions/verify", httpHandler.Middleware(httpHandler.VerifySession))

	fmt.Printf("listenning address %v\n", h.httpConfigs.Address)
	if err := http.ListenAndServe(h.httpConfigs.Address, nil); err != nil {
		fmt.Println(err)
	}
}

type standaloneServer struct {
	httpServer httpServer
	grpcServer grpcServer
}

func (s *standaloneServer) Start(ctx context.Context) {
	go func() {
		s.grpcServer.Start(ctx)
	}()

	s.httpServer.Start(ctx)
}

func NewStandaloneServer(
	configs configs.Config,
) StandaloneServer {
	return &standaloneServer{
		httpServer: httpServer{httpConfigs: configs.Http},
		grpcServer: grpcServer{grpcConfig: configs.Grpc},
	}
}
