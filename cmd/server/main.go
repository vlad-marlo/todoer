package main

import (
	"github.com/vlad-marlo/pkg/pgx"
	"github.com/vlad-marlo/todoer/internal/config"
	"github.com/vlad-marlo/todoer/internal/controller"
	"github.com/vlad-marlo/todoer/internal/controller/http"
	"github.com/vlad-marlo/todoer/internal/service"
	"github.com/vlad-marlo/todoer/internal/storage"
	pgxStorage "github.com/vlad-marlo/todoer/internal/storage/pgx"
	"github.com/vlad-marlo/todoer/pkg/pgx/client"
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
			http.New,
			config.NewServer,
			zap.NewProduction,
			fx.Annotate(client.New, fx.As(new(pgx.Client))),
			fx.Annotate(service.New, fx.As(new(controller.Service))),
			fx.Annotate(pgxStorage.New, fx.As(new(storage.Storage))),
			fx.Annotate(config.NewStorage, fx.As(new(client.Config))),
		),
		fx.Invoke(
			handleControllerLifecycle,
		),
	)
}

func handleControllerLifecycle(lc fx.Lifecycle, ctrl *http.Controller) {
	lc.Append(fx.Hook{
		OnStart: ctrl.Start,
		OnStop:  ctrl.Stop,
	})
}
