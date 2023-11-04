package main

import (
	"go.uber.org/fx"
	"go.uber.org/zap"
)

func main() {
	fx.New(NewServerOptions()).Run()
}

// NewServerOptions prepares fx options to start application.
func NewServerOptions() fx.Option {
	return fx.Options(
		fx.Provide(
			zap.NewProduction,
		),
	)
}
