//go:build wireinject
// +build wireinject

package handler

import (
	"github.com/ChatService/internal/configs"
	"github.com/ChatService/internal/dataaccess"
	"github.com/ChatService/internal/logic"
	"github.com/google/wire"
)

var WireSet = wire.NewSet(
	logic.WireSet,
	configs.Wireset,
	dataaccess.WireSet,
)

func GetAccountLogic(path configs.ConfigFilePath) (logic.Account, func(), error) {
	panic(wire.Build(WireSet))
}

func GetRelationshipLogic(path configs.ConfigFilePath) (logic.Relationship, func(), error) {
	panic(wire.Build(WireSet))
}
