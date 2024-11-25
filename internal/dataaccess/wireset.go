package dataaccess

import (
	"github.com/ChatService/internal/dataaccess/cache"
	"github.com/ChatService/internal/dataaccess/database"
	"github.com/ChatService/internal/dataaccess/services"
	"github.com/google/wire"
)

var WireSet = wire.NewSet(
	cache.WireSet,
	database.WireSet,
	services.WireSet,
)
