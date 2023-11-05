package memory

import (
	"sync"

	"github.com/google/uuid"
	"github.com/vlad-marlo/todoer/internal/model"
	"github.com/vlad-marlo/todoer/internal/storage"
	"go.uber.org/zap"
)

type Storage struct {
	mu    sync.Mutex
	tasks *repository
	log   *zap.Logger
}

// New is storage constructor.
func New(opts ...Option) storage.Storage {
	tasksRepository := &repository{
		data: make(map[uuid.UUID]*model.TaskDTO),
	}
	s := &Storage{
		tasks: tasksRepository,
	}

	var wg sync.WaitGroup

	for _, o := range opts {
		wg.Add(1)
		go applyAsync(&wg, s, o)
	}
	wg.Wait()

	return s
}

// Task return task repository to user.
func (s *Storage) Task() storage.TaskRepository {
	if s == nil {
		zap.L().Warn("storage: memory: Storage: Task(): nil reference")
		return nil
	}

	s.mu.Lock()
	tasks := s.tasks
	s.mu.Unlock()

	return tasks
}
