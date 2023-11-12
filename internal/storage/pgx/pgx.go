package pgx

import (
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/vlad-marlo/pkg/pgx"
	"github.com/vlad-marlo/todoer/internal/storage"
	"go.uber.org/zap"
)

var _ storage.Storage = (*Storage)(nil)

type Storage struct {
	log  *zap.Logger
	cli  *pgxpool.Pool
	task *TaskRepository
}

func (s *Storage) Task() storage.TaskRepository {
	return s.task
}

func New(log *zap.Logger, cli pgx.Client) (*Storage, error) {
	s := &Storage{
		log:  log.With(zap.String("layer", "storage")),
		cli:  cli.P(),
		task: newTask(log, cli.P()),
	}
	s.log.Info("initialized pgx storage")
	return s, nil
}
