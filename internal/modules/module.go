package modules

import (
	"github.com/mbrovarska/study-corner-user-service/internal/adapters/handler"
	"go.uber.org/fx"
)

var Module = fx.Options(
	fx.Provide(NewLogger),
	fx.Provide(handler.NewHealthHandler),
	fx.Provide(NewGinRouter),
	fx.Provide(handler.NewHTTPHandler),
	fx.Invoke(StartServer),
)