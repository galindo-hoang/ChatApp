package logic

import (
	"context"
	"github.com/ChatService/internal/dataaccess/database"
)

type Relationship interface {
	CreateNode(ctx context.Context, node AccountResponse) error
	FollowingPerson(ctx context.Context, from uint64, to uint64) error
	UnfollowingPerson(ctx context.Context, from uint64, to uint64) error
}

type relationship struct {
	relationshipDataAccessor database.RelationshipDataAccessor
}

func NewRelationship(
	relationshipDataAccessor database.RelationshipDataAccessor,
) Relationship {
	return &relationship{
		relationshipDataAccessor: relationshipDataAccessor,
	}
}

func (r *relationship) CreateNode(ctx context.Context, node AccountResponse) error {
	return r.relationshipDataAccessor.CreateNode(ctx, &database.Accounts{
		Id:          node.ID,
		Email:       node.Email,
		AccountName: node.AccountName,
	})
}

func (r *relationship) FollowingPerson(ctx context.Context, from uint64, to uint64) error {
	return r.relationshipDataAccessor.CreateRelationship(ctx, from, to)
}

func (r *relationship) UnfollowingPerson(ctx context.Context, from uint64, to uint64) error {
	return r.relationshipDataAccessor.RemoveRelationship(ctx, from, to)
}
