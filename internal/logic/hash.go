package logic

import (
	"context"
	"errors"
	"github.com/ChatService/internal/configs"
	"golang.org/x/crypto/bcrypt"
)

type Hash interface {
	Hash(ctx context.Context, data string) (string, error)
	isHashEqual(ctx context.Context, data, hash string) (bool, error)
}

type hash struct {
	authHash configs.Auth
}

func (h *hash) Hash(_ context.Context, data string) (string, error) {
	hashed, err := bcrypt.GenerateFromPassword([]byte(data), h.authHash.Hash.Cost)
	if err != nil {
		return "", err
	}
	return string(hashed), nil
}

func (h *hash) isHashEqual(_ context.Context, data, hash string) (bool, error) {
	if err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(data)); err != nil {
		if errors.Is(err, bcrypt.ErrMismatchedHashAndPassword) {
			return false, nil
		}
		return false, err
	}
	return true, nil
}

func NewHash(authHash configs.Auth) Hash {
	return &hash{
		authHash: authHash,
	}
}
