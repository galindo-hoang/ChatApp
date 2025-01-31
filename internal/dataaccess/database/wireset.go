package database

import "github.com/google/wire"

var WireSet = wire.NewSet(
	InitializeAccountDataAccessor,
	InitializeMessageDataAccessor,
	InitializeMysqlDataAccessor,
	InitializeDB,
	InitializeGorm,
	InitializeGraphDB,
)
