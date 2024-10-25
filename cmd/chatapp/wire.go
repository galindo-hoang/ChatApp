//go:build wireinject
// +build wireinject

package main

import "github.com/google/wire"

func Initialize() (C, error) {
	panic(wire.Build(
		wire.NewSet(
			GetContentA,
			ContentB,
			ContentA,
			GetContentB,
			wire.FieldsOf(new(B), "contentC"),
		),
	))
}
