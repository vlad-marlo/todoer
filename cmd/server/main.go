package main

import (
	"go.uber.org/fx"
	"go.uber.org/zap"
)

func main() {
	fx.New(NewServerOptions()).Run()
}

func NewServerOptions() fx.Option {
	return fx.Options(
		fx.Provide(
			zap.NewProduction,
		),
	)
}
