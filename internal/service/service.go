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

func New(logger *zap.Logger) (*Service, error) {
	return &Service{
		log:     logger,
		storage: nil,
	}, nil
}
