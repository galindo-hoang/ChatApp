package wiring

import (
	"github.com/ChatService/internal/configs"
	"github.com/ChatService/internal/dataaccess"
	"github.com/ChatService/internal/logic"
	"github.com/google/wire"
)

var WireSet = wire.NewSet(
	dataaccess.WireSet,
	configs.Wireset,
	logic.WireSet,
)
