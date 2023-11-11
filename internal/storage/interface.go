package storage

import (
	"context"
	"github.com/vlad-marlo/todoer/internal/model"
)

// Storage is interface of storage, which can be used in application.
type Storage interface {
	Task() TaskRepository
}

// TaskRepository is interface of object, which
// is storing and returning tasks.
type TaskRepository interface {
	Create(ctx context.Context, task *model.TaskDTO) error
}
