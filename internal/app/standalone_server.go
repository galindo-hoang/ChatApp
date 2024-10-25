package app

import "github.com/ChatService/internal/dataaccess/database"

type StandaloneServer struct {
	accountAccessor database.AccountDataAccessor
	messageAccessor database.MessageDataAccessor
}

func NewStandaloneServer(
	account database.AccountDataAccessor,
	message database.MessageDataAccessor,
) *StandaloneServer {
	return &StandaloneServer{
		accountAccessor: account,
		messageAccessor: message,
	}
}
