package handler

import (
	"github.com/ChatService/internal/handler/http"
	"github.com/google/wire"
)

var WireSet = wire.NewSet(
	http.WireSet,
)
