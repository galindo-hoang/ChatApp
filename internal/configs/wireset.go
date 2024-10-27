package configs

import "github.com/google/wire"

var Wireset = wire.NewSet(
	NewConfig,
	wire.FieldsOf(new(Config), "Database"),
	wire.FieldsOf(new(Config), "Auth"),
)
