package wiring

import (
	"github.com/ChatService/internal/configs"
	"github.com/ChatService/internal/dataaccess/database"
	"github.com/google/wire"
)

var WireSet = wire.NewSet(
	database.WireSet,
	configs.Wireset,
)
