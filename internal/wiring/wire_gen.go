// Code generated by Wire. DO NOT EDIT.

//go:generate go run -mod=mod github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package wiring

import (
	"github.com/ChatService/internal/app"
	"github.com/ChatService/internal/configs"
	"github.com/ChatService/internal/dataaccess"
	"github.com/ChatService/internal/dataaccess/cache"
	"github.com/ChatService/internal/dataaccess/database"
	"github.com/ChatService/internal/handler"
	"github.com/ChatService/internal/handler/http"
	"github.com/ChatService/internal/logic"
	"github.com/google/wire"
)

// Injectors from wire.go:

func InitializeStandaloneServer(path configs.ConfigFilePath) (app.StandaloneServer, func(), error) {
	config, err := configs.NewConfig(path)
	if err != nil {
		return nil, nil, err
	}
	configsDatabase := config.Database
	db, cleanup, err := database.InitializeAndMigrateUpDB(configsDatabase)
	if err != nil {
		return nil, nil, err
	}
	gormDB, err := database.InitializeGorm(db, configsDatabase)
	if err != nil {
		cleanup()
		return nil, nil, err
	}
	auth := config.Auth
	hash := logic.NewHash(auth)
	token, err := logic.NewToken(auth)
	if err != nil {
		cleanup()
		return nil, nil, err
	}
	takenAccountEmail := cache.NewTakenAccountEmail()
	accountDataAccessor := database.InitializeAccountDataAccessor(gormDB)
	account := logic.NewAccount(gormDB, hash, token, takenAccountEmail, accountDataAccessor)
	httpServer := http.NewHttpServer(account, config)
	standaloneServer := app.NewStandaloneServer(httpServer)
	return standaloneServer, func() {
		cleanup()
	}, nil
}

// wire.go:

var WireSet = wire.NewSet(app.WireSet, logic.WireSet, handler.WireSet, configs.Wireset, dataaccess.WireSet)