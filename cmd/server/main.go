package main

import (
	"github.com/vlad-marlo/todoer/internal/config"
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
			StorageProvider,
		),
	)
}

func StorageProvider(log *zap.Logger, cfg *config.Storage) (storage.Storage, error) {
	return memory.New(memory.WithZapLogger(log)), nil
}
