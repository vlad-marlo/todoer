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
	Count(ctx context.Context) (int, error)
	Paginate(ctx context.Context, offset int, limit int) ([]model.TaskDTO, error)
	PaginateFilter(ctx context.Context, offset int, limit int, task string) ([]model.TaskDTO, error)
	PaginateWithStatus(ctx context.Context, offset int, limit int, status model.Status) ([]model.TaskDTO, error)
	PaginateWithStatuses(ctx context.Context, offset int, limit int, status1, status2 model.Status) ([]model.TaskDTO, error)
	PaginateWithoutStatus(ctx context.Context, offset int, limit int, ignored model.Status) ([]model.TaskDTO, error)
	PaginateFilterWithStatus(ctx context.Context, offset int, limit int, task string, status model.Status) ([]model.TaskDTO, error)
	PaginateFilterWithStatuses(ctx context.Context, offset int, limit int, task string, status1, status2 model.Status) ([]model.TaskDTO, error)
	PaginateFilterWithoutStatus(ctx context.Context, offset int, limit int, task string, ignored model.Status) ([]model.TaskDTO, error)
}
