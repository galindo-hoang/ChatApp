//go:build wireinject
// +build wireinject

package wiring

import (
	"github.com/ChatService/internal/app"
	"github.com/ChatService/internal/configs"
	"github.com/ChatService/internal/dataaccess"
	"github.com/ChatService/internal/handler"
	"github.com/ChatService/internal/logic"
	"github.com/google/wire"
)

var WireSet = wire.NewSet(
	app.WireSet,
	logic.WireSet,
	handler.WireSet,
	configs.Wireset,
	dataaccess.WireSet,
)

func InitializeStandaloneServer(path configs.ConfigFilePath) (app.StandaloneServer, func(), error) {
	panic(wire.Build(WireSet))
}
