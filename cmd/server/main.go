package main

import (
	"github.com/vlad-marlo/pkg/pgx"
	"github.com/vlad-marlo/pkg/pgx/client"
	"github.com/vlad-marlo/todoer/internal/config"
	"github.com/vlad-marlo/todoer/internal/controller"
	"github.com/vlad-marlo/todoer/internal/controller/http"
	"github.com/vlad-marlo/todoer/internal/service"
	"github.com/vlad-marlo/todoer/internal/storage"
	"github.com/vlad-marlo/todoer/internal/storage/memory"
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
			config.NewStorage,
			config.NewServer,
			http.New,
			fx.Annotate(service.New, fx.As(new(controller.Service))),
			StorageProvider,
			fx.Annotate(client.New, fx.As(new(pgx.Client))),
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

func StorageProvider(log *zap.Logger, cfg *config.Storage) (storage.Storage, error) {
	return memory.New(memory.WithZapLogger(log)), nil
}
