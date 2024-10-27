package logic

import (
	"context"
	"errors"
	"fmt"
	"github.com/ChatService/internal/configs"
	"github.com/golang-jwt/jwt/v5"
	"time"
)

type Token interface {
	GetToken(ctx context.Context, accountID uint64) (string, time.Time, error)
	GetAccountIDAndExpireTime(ctx context.Context, token string) (uint64, time.Time, error)
}

type token struct {
	authConfig configs.Auth
	expiresIn  time.Duration
	signedKey  string
	//privateKey       *rsa.PrivateKey
	//tokenPublicKeyID uint64
}

func NewToken(
	authConfig configs.Auth,
) (Token, error) {
	var expiresIn = authConfig.Token.GetExpiresInDuration()
	var signedKey = authConfig.Token.SignedKey
	return &token{
		authConfig: authConfig,
		expiresIn:  expiresIn,
		signedKey:  signedKey,
	}, nil
}

// token with symmetric

func (t *token) GetToken(ctx context.Context, accountID uint64) (string, time.Time, error) {
	fmt.Println("get token")

	expireTime := time.Now().Add(t.authConfig.Token.GetExpiresInDuration())
	token := jwt.NewWithClaims(jwt.SigningMethodHS512, jwt.MapClaims{
		"sub": accountID,
		"exp": expireTime.Unix(),
		//"kid": t.tokenPublicKeyID,
	})

	tokenString, err := token.SignedString([]byte(t.signedKey))
	if err != nil {
		fmt.Println(err.Error())
		return "", time.Time{}, err
	}
	return tokenString, expireTime, nil
}

func (t *token) GetAccountIDAndExpireTime(ctx context.Context, token string) (uint64, time.Time, error) {
	parsedToken, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		return []byte(t.signedKey), nil
	})
	if err != nil {
		return 0, time.Time{}, err
	}

	if !parsedToken.Valid {
		return 0, time.Time{}, errors.New("invalid token")
	}
	claims, ok := parsedToken.Claims.(jwt.MapClaims)
	if !ok {
		fmt.Println("can't get token claims")
		return 0, time.Time{}, errors.New("invalid token")
	}

	accountId, ok := claims["sub"].(float64)
	if !ok {
		fmt.Println("can't get token's claims")
		return 0, time.Time{}, errors.New("can't get token's claims")
	}

	expireTimeUnix, ok := claims["exp"].(float64)
	if !ok {
		fmt.Println("can't get token's exp claims")
		return 0, time.Time{}, errors.New("can't get token's exp claims")
	}

	return uint64(accountId), time.Unix(int64(expireTimeUnix), 0), nil
}
