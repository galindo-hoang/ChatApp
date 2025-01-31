package app

import (
	"context"
	"fmt"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/ChatService/internal/configs"
	pb "github.com/ChatService/internal/generated/chat_app/v1"
	grpcHandler "github.com/ChatService/internal/handler/grpc"
	httpHandler "github.com/ChatService/internal/handler/http"
	"github.com/confluentinc/confluent-kafka-go/kafka"
	"google.golang.org/grpc"
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

	http.HandleFunc("/v1/friends/add", httpHandler.Middleware(httpHandler.AddFriend))
	http.HandleFunc("/v1/friends/remove", httpHandler.Middleware(httpHandler.RemoveFriend))
	http.HandleFunc("/v1/friends/decline", httpHandler.Middleware(httpHandler.DeclineFriend))
	http.HandleFunc("/v1/friends/accept", httpHandler.Middleware(httpHandler.AcceptRequest))
	http.HandleFunc("/v1/friends/requests", httpHandler.Middleware(httpHandler.GetListRequests))
	http.HandleFunc("/v1/friends/pendings", httpHandler.Middleware(httpHandler.GetListPendings))

	fmt.Printf("listenning address %v\n", h.httpConfigs.Address)
	if err := http.ListenAndServe(h.httpConfigs.Address, nil); err != nil {
		fmt.Println(err)
	}
}

type kafkaServer struct {
	kafkaConfigs configs.Kafka
}

func (k *kafkaServer) Start(ctx context.Context) {
	c, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers": fmt.Sprintf("%v:%v", k.kafkaConfigs.Host, k.kafkaConfigs.Port),

		// Fixed properties
		"group.id":          k.kafkaConfigs.GroupId,
		"auto.offset.reset": "earliest",
	})

	if err != nil {
		fmt.Printf("Failed to create consumer: %s\n", err)
		os.Exit(1)
	}

	err = c.SubscribeTopics([]string{k.kafkaConfigs.Topic}, nil)
	if err != nil {
		fmt.Printf("Failed to subscribe: %s\n", err)
	}

	sigchan := make(chan os.Signal, 1)
	signal.Notify(sigchan, syscall.SIGINT, syscall.SIGTERM)

	run := true

	for run {
		select {

		case sig := <-sigchan:
			fmt.Printf("Caught signal %v: terminating \n", sig)
			run = false
		default:
			ev, err := c.ReadMessage(100 * time.Microsecond)
			if err != nil {
				continue
			}
			fmt.Printf("Consumed event from topic %s: key = %-10s value = %s\n",
				*ev.TopicPartition.Topic, string(ev.Key), string(ev.Value))
		}
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
