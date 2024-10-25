//go:build wireinject
// +build wireinject

package database

import (
	"github.com/ChatService/internal/configs"
	"github.com/google/wire"
)

var wireSetInternal = wire.NewSet(
	WireSet,
	configs.Wireset,
)

func getAccountAccessor(
	path configs.ConfigFilePath,
) (AccountDataAccessor, func(), error) {
	panic(wire.Build(wireSetInternal))
}
