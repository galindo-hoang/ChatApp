package dataaccess

import (
	"github.com/ChatService/internal/dataaccess/cache"
	"github.com/ChatService/internal/dataaccess/database"
	"github.com/google/wire"
)

var WireSet = wire.NewSet(
	cache.WireSet,
	database.WireSet,
)
