package service

import (
	"github.com/vlad-marlo/todoer/internal/controller"
	"github.com/vlad-marlo/todoer/internal/storage"
	"go.uber.org/zap"
)

var _ controller.Service = (*Service)(nil)

type Service struct {
	log     *zap.Logger
	storage storage.Storage
}

func New(logger *zap.Logger, repo storage.Storage) (*Service, error) {
	logger.Info("initialized service")
	return &Service{
		log:     logger.With(zap.String("layer", "service")),
		storage: repo,
	}, nil
}
