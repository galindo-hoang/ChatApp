//go:build wireinject
// +build wireinject

package logic

import (
	"github.com/ChatService/internal/configs"
	"github.com/ChatService/internal/dataaccess"
	"github.com/google/wire"
)

func initializeToken(filePath configs.ConfigFilePath) (Token, error) {
	panic(wire.Build(WireSet, configs.Wireset))
}

func initializeAccount(filePath configs.ConfigFilePath) (Account, func(), error) {
	panic(wire.Build(WireSet, configs.Wireset, dataaccess.WireSet))
}
