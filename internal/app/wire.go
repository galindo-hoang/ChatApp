//go:build wireinject
// +build wireinject

package app

import (
	"github.com/ChatService/internal/configs"
	"github.com/google/wire"
)

var WireSet = wire.NewSet(
	configs.Wireset,
	NewStandaloneServer,
)

func InitializeStandaloneServer(path configs.ConfigFilePath) (StandaloneServer, error) {
	panic(wire.Build(WireSet))
}
