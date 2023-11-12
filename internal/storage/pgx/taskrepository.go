package pgx

import (
	"context"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/vlad-marlo/todoer/internal/model"
	"go.uber.org/zap"
)

type TaskRepository struct {
	log  *zap.Logger
	pool *pgxpool.Pool
}

func newTask(log *zap.Logger, pool *pgxpool.Pool) *TaskRepository {
	repo := &TaskRepository{
		log:  log.With(zap.String("entity", "task repository")),
		pool: pool,
	}
	repo.log.Info("initialized task repository")
	return repo
}

func (t *TaskRepository) Get(ctx context.Context, id uuid.UUID) (*model.TaskDTO, error) {
	const query = `SELECT id, value, created_at, status
FROM tasks
where id = $1;`

	var (
		task   = new(model.TaskDTO)
		status int
	)

	if err := t.pool.QueryRow(ctx, query, id).Scan(&task.ID, &task.Value, &task.CreatedAt, &status); err != nil {
		return nil, err
	}
	task.Status = model.StatusFromInt(status)
	return task, nil
}

func (t *TaskRepository) Count(ctx context.Context) (int, error) {
	const query = `SELECT COUNT(*)
FROM tasks;`

	var result int

	if err := t.pool.QueryRow(ctx, query).Scan(&result); err != nil {
		return 0, err
	}
	return result, nil
}

func (t *TaskRepository) Create(ctx context.Context, task *model.TaskDTO) error {
	const query = `INSERT INTO tasks(id, value, created_at, status)
VALUES ($1, $2, $3, $4);`

	if _, err := t.pool.Exec(ctx, query, task.ID, task.Value, task.CreatedAt, task.Status.Int()); err != nil {
		return err
	}

	return nil
}
func (t *TaskRepository) Update(ctx context.Context, id uuid.UUID, value string, status model.Status) error {
	const query = `UPDATE tasks
SET value  = $1,
    status = $2
WHERE id = $3;`

	if _, err := t.pool.Exec(ctx, query, value, status.Int(), id); err != nil {
		return err
	}
	return nil
}

func (t *TaskRepository) Paginate(ctx context.Context, offset int, limit int) ([]model.TaskDTO, error) {
	//TODO implement me
	panic("implement me")
}

func (t *TaskRepository) CreateMany(ctx context.Context, tasks []model.TaskDTO) error {
	//TODO implement me
	panic("implement me")
}

func (t *TaskRepository) ChangeStatus(ctx context.Context, id uuid.UUID, status model.Status) error {
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
