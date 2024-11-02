package logic

import (
	"context"
	"errors"
	"fmt"
	"github.com/ChatService/internal/dataaccess/cache"
	"github.com/ChatService/internal/dataaccess/database"
	"gorm.io/gorm"
	"time"
)

type Account interface {
	CreateAccount(ctx context.Context, params CreateAccountParams) (CreateAccountResponse, error)
	CreateSession(ctx context.Context, params CreateSessionParams) (CreateSessionResponse, error)
	ValidateSession(ctx context.Context, session string) error
}

type account struct {
	accountDataAccessor database.AccountDataAccessor
	takenAccountCache   cache.TakenAccountEmail
	db                  *gorm.DB
	hashLogic           Hash
	tokenLogic          Token
}

func NewAccount(
	db *gorm.DB,
	hashLogic Hash,
	tokenLogic Token,
	takenAccountCache cache.TakenAccountEmail,
	accountDataAccessor database.AccountDataAccessor,
) Account {
	return &account{
		db:                  db,
		hashLogic:           hashLogic,
		tokenLogic:          tokenLogic,
		takenAccountCache:   takenAccountCache,
		accountDataAccessor: accountDataAccessor,
	}
}

func (a *account) isAccountEmailTaken(ctx context.Context, email string) (bool, error) {
	fmt.Println("isAccountEmailTaken checking...")

	isExist, err := a.takenAccountCache.Has(ctx, email)
	if err != nil {
		fmt.Println(err.Error())
	} else if isExist {
		return true, nil
	}

	res, err := a.accountDataAccessor.GetAccountByEmail(ctx, email)
	if err != nil {
		return false, err
	}
	if res == nil {
		return false, nil
	}

	err = a.takenAccountCache.Add(ctx, email)
	if err != nil {
		fmt.Println(err.Error())
	}

	return true, nil
}

func (a *account) CreateAccount(ctx context.Context, params CreateAccountParams) (CreateAccountResponse, error) {
	isExist, err := a.isAccountEmailTaken(ctx, params.Email)
	if err != nil {
		return CreateAccountResponse{}, errors.New(fmt.Sprintf("fail to check account email taken: %s", err.Error()))
	}
	if isExist {
		return CreateAccountResponse{}, errors.New(fmt.Sprintf("account email taken: %s", params.Email))
	}

	var accountId uint64
	err = a.db.Transaction(func(tx *gorm.DB) error {
		hashPassword, err := a.hashLogic.Hash(ctx, params.Password)
		if err != nil {
			return err
		}
		account, err := a.accountDataAccessor.CreateAccount(ctx, &database.Accounts{
			AccountName: params.AccountName,
			Email:       params.Email,
			Password:    hashPassword,
		})
		accountId = account.Id

		return nil
	})
	if err != nil {
		return CreateAccountResponse{}, err
	}

	return CreateAccountResponse{
		ID:    accountId,
		Email: params.Email,
	}, nil
}

func (a *account) CreateSession(ctx context.Context, params CreateSessionParams) (CreateSessionResponse, error) {
	fmt.Println("CreateSession...")
	account, err := a.accountDataAccessor.GetAccountByEmail(ctx, params.Email)
	if err != nil {
		return CreateSessionResponse{}, err
	}

	if account == nil {
		return CreateSessionResponse{}, errors.New(fmt.Sprint("account or password are invalid"))
	}

	isEqual, err := a.hashLogic.isHashEqual(ctx, params.Password, account.Password)
	if err != nil {
		return CreateSessionResponse{}, err
	}
	if !isEqual {
		return CreateSessionResponse{}, errors.New(fmt.Sprint("account or password is invalid"))
	}

	token, _, err := a.tokenLogic.GetToken(ctx, account.Id)
	if err != nil {
		return CreateSessionResponse{}, err
	}

	return CreateSessionResponse{
		Account: AccountResponse{
			ID:          account.Id,
			Email:       account.Email,
			AccountName: account.AccountName,
		},
		AccessToken: token,
	}, nil

}

func (a *account) ValidateSession(ctx context.Context, session string) error {
	fmt.Println("ValidateSession checking...")

	id, ti, err := a.tokenLogic.GetAccountIDAndExpireTime(ctx, session)
	if err != nil {
		return err
	}

	if ti.Before(time.Now()) {
		return errors.New("token is invalid")
	}

	account, err := a.accountDataAccessor.GetAccountByID(ctx, id)
	if err != nil || account == nil {
		return errors.New("token is invalid")
	}

	return nil
}
