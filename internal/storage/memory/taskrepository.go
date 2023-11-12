package memory

import (
	"context"
	"sync"

	"github.com/google/uuid"
	"github.com/vlad-marlo/todoer/internal/model"
	"go.uber.org/zap"
)

type repository struct {
	mu   sync.Mutex
	log  *zap.Logger
	data map[uuid.UUID]*model.TaskDTO
}

func (r *repository) Get(ctx context.Context, id string) (*model.TaskDTO, error) {
	//TODO implement me
	panic("implement me")
}

func (r *repository) Count(ctx context.Context) (int, error) {
	//TODO implement me
	panic("implement me")
}

func (r *repository) Create(ctx context.Context, task *model.TaskDTO) error {
	//TODO implement me
	panic("implement me")
}

func (r *repository) Update(ctx context.Context, id, value string, status model.Status) error {
	//TODO implement me
	panic("implement me")
}

func (r *repository) Paginate(ctx context.Context, offset int, limit int) ([]model.TaskDTO, error) {
	//TODO implement me
	panic("implement me")
}

func (r *repository) CreateMany(ctx context.Context, tasks []model.TaskDTO) error {
	//TODO implement me
	panic("implement me")
}

func (r *repository) ChangeStatus(ctx context.Context, id string, status model.Status) error {
	//TODO implement me
	panic("implement me")
}

func (r *repository) PaginateFilter(ctx context.Context, offset int, limit int, task string) ([]model.TaskDTO, error) {
	//TODO implement me
	panic("implement me")
}

func (r *repository) PaginateWithStatus(ctx context.Context, offset int, limit int, status model.Status) ([]model.TaskDTO, error) {
	//TODO implement me
	panic("implement me")
}

func (r *repository) PaginateWithStatuses(ctx context.Context, offset int, limit int, status1, status2 model.Status) ([]model.TaskDTO, error) {
	//TODO implement me
	panic("implement me")
}

func (r *repository) PaginateWithoutStatus(ctx context.Context, offset int, limit int, ignored model.Status) ([]model.TaskDTO, error) {
	//TODO implement me
	panic("implement me")
}

func (r *repository) PaginateFilterWithStatus(ctx context.Context, offset int, limit int, task string, status model.Status) ([]model.TaskDTO, error) {
	//TODO implement me
	panic("implement me")
}

func (r *repository) PaginateFilterWithStatuses(ctx context.Context, offset int, limit int, task string, status1, status2 model.Status) ([]model.TaskDTO, error) {
	//TODO implement me
	panic("implement me")
}

func (r *repository) PaginateFilterWithoutStatus(ctx context.Context, offset int, limit int, task string, ignored model.Status) ([]model.TaskDTO, error) {
	//TODO implement me
	panic("implement me")
}
