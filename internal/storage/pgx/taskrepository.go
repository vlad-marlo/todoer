package pgx

import (
	"context"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/vlad-marlo/todoer/internal/model"
	"go.uber.org/zap"
)

type TaskRepository struct {
	log  *zap.Logger
	pool *pgxpool.Pool
}

func newTask(log *zap.Logger, pool *pgxpool.Pool) *TaskRepository {
	return &TaskRepository{
		log:  log.With(zap.String("entity", "task repository")),
		pool: pool,
	}
}

func (t *TaskRepository) Get(ctx context.Context, id string) (*model.TaskDTO, error) {
	//TODO implement me
	panic("implement me")
}

func (t *TaskRepository) Count(ctx context.Context) (int, error) {
	//TODO implement me
	panic("implement me")
}

func (t *TaskRepository) Create(ctx context.Context, task *model.TaskDTO) error {
	//TODO implement me

	panic("implement me")
}
func (t *TaskRepository) Update(ctx context.Context, id, value string, status model.Status) error {
	//TODO implement me
	panic("implement me")
}

func (t *TaskRepository) Paginate(ctx context.Context, offset int, limit int) ([]model.TaskDTO, error) {
	//TODO implement me
	panic("implement me")
}

func (t *TaskRepository) CreateMany(ctx context.Context, tasks []model.TaskDTO) error {
	//TODO implement me
	panic("implement me")
}

func (t *TaskRepository) ChangeStatus(ctx context.Context, id string, status model.Status) error {
	//TODO implement me
	panic("implement me")
}

func (t *TaskRepository) PaginateFilter(ctx context.Context, offset int, limit int, task string) ([]model.TaskDTO, error) {
	//TODO implement me
	panic("implement me")
}

func (t *TaskRepository) PaginateWithStatus(ctx context.Context, offset int, limit int, status model.Status) ([]model.TaskDTO, error) {
	//TODO implement me
	panic("implement me")
}

func (t *TaskRepository) PaginateWithStatuses(ctx context.Context, offset int, limit int, status1, status2 model.Status) ([]model.TaskDTO, error) {
	//TODO implement me
	panic("implement me")
}

func (t *TaskRepository) PaginateWithoutStatus(ctx context.Context, offset int, limit int, ignored model.Status) ([]model.TaskDTO, error) {
	//TODO implement me
	panic("implement me")
}

func (t *TaskRepository) PaginateFilterWithStatus(ctx context.Context, offset int, limit int, task string, status model.Status) ([]model.TaskDTO, error) {
	//TODO implement me
	panic("implement me")
}

func (t *TaskRepository) PaginateFilterWithStatuses(ctx context.Context, offset int, limit int, task string, status1, status2 model.Status) ([]model.TaskDTO, error) {
	//TODO implement me
	panic("implement me")
}

func (t *TaskRepository) PaginateFilterWithoutStatus(ctx context.Context, offset int, limit int, task string, ignored model.Status) ([]model.TaskDTO, error) {
	//TODO implement me
	panic("implement me")
}
