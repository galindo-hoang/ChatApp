package grpc_chat

import (
	"context"
	"fmt"
	pb "github.com/ChatService/internal/generated/chat_app/v1"
	"github.com/ChatService/internal/handler"
	"github.com/ChatService/internal/logic"
)

type RelationshipHandler struct {
	pb.UnimplementedRelationshipServiceServer
}

func (r *RelationshipHandler) CreateNode(ctx context.Context, request *pb.CreateAccountRequest) (*pb.CreateAccountResponse, error) {
	fmt.Printf("grpc create node: %v\n", request)

	relationshipLogic, clearLogic, err := handler.GetRelationshipLogic("")
	defer clearLogic()
	if err != nil {
		return nil, err
	}

	err = relationshipLogic.CreateNode(ctx, logic.AccountResponse{
		ID:          request.AccountId,
		AccountName: request.AccountName,
		Email:       request.Email,
	})

	if err != nil {
		return &pb.CreateAccountResponse{
			IsSuccess: false,
			Message:   err.Error(),
		}, err

	}
	return &pb.CreateAccountResponse{IsSuccess: true}, nil
}

func (r *RelationshipHandler) FollowingPerson(ctx context.Context, request *pb.MakeRelationshipRequest) (*pb.MakeRelationshipResponse, error) {
	fmt.Printf("grpc FollowingPerson: %v\n", request)

	relationshipLogic, clearLogic, err := handler.GetRelationshipLogic("")
	defer clearLogic()
	if err != nil {
		return nil, err
	}

	err = relationshipLogic.FollowingPerson(ctx, request.From, request.To)

	if err != nil {
		return &pb.MakeRelationshipResponse{
			IsSuccess: false,
			Message:   err.Error(),
		}, err

	}
	return &pb.MakeRelationshipResponse{IsSuccess: true}, nil
}

func (r *RelationshipHandler) UnFollowingPerson(ctx context.Context, request *pb.MakeRelationshipRequest) (*pb.MakeRelationshipResponse, error) {
	fmt.Printf("grpc UnFollowingPerson: %v\n", request)

	relationshipLogic, clearLogic, err := handler.GetRelationshipLogic("")
	defer clearLogic()
	if err != nil {
		return nil, err
	}

	err = relationshipLogic.UnfollowingPerson(ctx, request.From, request.To)

	if err != nil {
		return &pb.MakeRelationshipResponse{
			IsSuccess: false,
			Message:   err.Error(),
		}, err

	}

	return &pb.MakeRelationshipResponse{IsSuccess: true}, nil
}
