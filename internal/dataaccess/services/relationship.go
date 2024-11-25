package services

import (
	"fmt"
	"github.com/ChatService/internal/configs"
	pb "github.com/ChatService/internal/generated/chat_app/v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func InitializeRelationshipClient(grpcConfig configs.GRPC /*, logger *zap.Logger*/) (pb.RelationshipServiceClient, func(), error) {
	fmt.Println("InitializeRelationshipClient initialize...")
	conn, err := grpc.NewClient(fmt.Sprintf("%v:%v", grpcConfig.Host, grpcConfig.Port), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		fmt.Printf("Failed to create gRPC connection: %v\n", err)
		return nil, nil, err
	}

	cleanup := func() {
		fmt.Println("InitializeRelationshipClient cancel...")
		conn.Close()
	}

	c := pb.NewRelationshipServiceClient(conn)
	return c, cleanup, nil
}
