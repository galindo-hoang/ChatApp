package logic

import (
	"context"
	"fmt"

	"github.com/ChatService/internal/dataaccess/database"
)

type Relationship interface {
	RemoveFriend(context.Context, uint64, uint64) error
	RequestFriend(context.Context, uint64, uint64) error
	AcceptFriend(context.Context, uint64, uint64) error
	GetListFriends(context.Context, uint64) ([]*AccountResponse, error)
	GetRequestFriends(context.Context, uint64) ([]*AccountResponse, error)
	GetPendingFriends(context.Context, uint64) ([]*AccountResponse, error)
}

type relationship struct {
	relationshipDataAccessor database.RelationshipDataAccessor
	accountDataAccessor      database.AccountDataAccessor
}

func NewRelationship(
	relationshipDataAccessor database.RelationshipDataAccessor,
	accountDataAccessor database.AccountDataAccessor,
) Relationship {
	return &relationship{
		relationshipDataAccessor: relationshipDataAccessor,
		accountDataAccessor: accountDataAccessor,
	}
}

func (r *relationship) RemoveFriend(ctx context.Context, from uint64, to uint64) error {
	return r.relationshipDataAccessor.RemoveRelationship(ctx, from, to)
}
func (r *relationship) RequestFriend(ctx context.Context, from uint64, to uint64) error {
	return r.relationshipDataAccessor.CreateRelationship(ctx, from, to)
}
func (r *relationship) AcceptFriend(ctx context.Context, from uint64, to uint64) error {
	return r.relationshipDataAccessor.UpdateRelationship(ctx, from, to)
}
func (r *relationship) GetListFriends(ctx context.Context, from uint64) ([]*AccountResponse, error) {
	ids, err := r.relationshipDataAccessor.GetListFriends(ctx, from)
	if err != nil {
		return nil, err
	}

	var result []*AccountResponse
	for _, id := range ids {
		account, err := r.accountDataAccessor.GetAccountByID(ctx, *id)
		if err != nil {
			fmt.Printf("error get account %v", account)
		}
		var parsedAccount = &AccountResponse{}
		parsedAccount.AccountName = account.AccountName
		parsedAccount.Email = account.Email
		parsedAccount.ID = account.Id
		result = append(result, parsedAccount)
	}
	return result, nil
}
func (r *relationship) GetRequestFriends(ctx context.Context, from uint64) ([]*AccountResponse, error) {
	ids, err := r.relationshipDataAccessor.GetListRequests(ctx, from)
	if err != nil {
		return nil, err
	}

	var result []*AccountResponse
	for _, id := range ids {
		account, err := r.accountDataAccessor.GetAccountByID(ctx, *id)
		if err != nil {
			fmt.Printf("error get account %v", account)
		}
		var parsedAccount = &AccountResponse{}
		parsedAccount.AccountName = account.AccountName
		parsedAccount.Email = account.Email
		parsedAccount.ID = account.Id
		result = append(result, parsedAccount)
	}
	return result, nil
}
func (r *relationship) GetPendingFriends(ctx context.Context, from uint64) ([]*AccountResponse, error) {
	ids, err := r.relationshipDataAccessor.GetListPendingRequesters(ctx, from)
	if err != nil {
		return nil, err
	}

	var result []*AccountResponse
	for _, id := range ids {
		account, err := r.accountDataAccessor.GetAccountByID(ctx, *id)
		if err != nil {
			fmt.Printf("error get account %v", account)
		}
		var parsedAccount = &AccountResponse{}
		parsedAccount.AccountName = account.AccountName
		parsedAccount.Email = account.Email
		parsedAccount.ID = account.Id
		result = append(result, parsedAccount)
	}
	return result, nil
}
