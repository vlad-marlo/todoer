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
WHERE id = $1;`

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

func (t *TaskRepository) getManyTasksByQuery(ctx context.Context, query string, args ...any) ([]model.TaskDTO, error) {
	var res []model.TaskDTO

	rows, err := t.pool.Query(ctx, query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var (
			task   model.TaskDTO
			status int
		)
		if err = rows.Scan(&task.ID, &task.Value, &task.CreatedAt, &status); err != nil {
			return nil, err
		}
		task.Status = model.StatusFromInt(status)
		res = append(res, task)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}
	return res, nil
}

func (t *TaskRepository) Paginate(ctx context.Context, offset int, limit int) ([]model.TaskDTO, error) {
	const query = `SELECT id, value, created_at, status
FROM tasks
ORDER BY created_at DESC
OFFSET $1 LIMIT $2;`
	return t.getManyTasksByQuery(ctx, query, offset, limit)
}

func (t *TaskRepository) PaginateFilter(ctx context.Context, offset int, limit int, task string) ([]model.TaskDTO, error) {
	const query = `SELECT id, value, created_at, status
FROM tasks
WHERE value LIKE CONCAT('%%',$1::text,'%%')
ORDER BY created_at DESC
OFFSET $2 LIMIT $3;`
	return t.getManyTasksByQuery(ctx, query, task, offset, limit)
}

func (t *TaskRepository) CreateMany(context.Context, []model.TaskDTO) error {
	//TODO implement me
	panic("implement me")
}

func (t *TaskRepository) ChangeStatus(ctx context.Context, id uuid.UUID, status model.Status) error {
	const query = `UPDATE tasks
SET status = $1
WHERE id = $2;`
	if _, err := t.pool.Exec(ctx, query, status.Int(), id); err != nil {
		return err
	}
	return nil
}

func (t *TaskRepository) PaginateWithStatus(ctx context.Context, offset int, limit int, status model.Status) ([]model.TaskDTO, error) {
	t.log.Info("paginate with status", zap.Int("offset", offset), zap.Int("limit", limit), zap.String("status", status.String()))
	const query = `SELECT id, value, created_at, status
FROM tasks
WHERE status = $1
ORDER BY created_at DESC
OFFSET $2 LIMIT $3;`
	return t.getManyTasksByQuery(ctx, query, status.Int(), offset, limit)
}

func (t *TaskRepository) PaginateWithStatuses(ctx context.Context, offset int, limit int, status1, status2 model.Status) ([]model.TaskDTO, error) {
	t.log.Info("paginate with statuses", zap.Int("offset", offset), zap.Int("limit", limit), zap.String("status1", status1.String()), zap.String("status2", status1.String()))
	const query = `SELECT id, value, created_at, status
FROM tasks
WHERE status = $1 OR status = $2
ORDER BY created_at DESC
OFFSET $3 LIMIT $4;`
	return t.getManyTasksByQuery(ctx, query, status1.Int(), status2.Int(), offset, limit)
}

func (t *TaskRepository) PaginateWithoutStatus(ctx context.Context, offset int, limit int, ignored model.Status) ([]model.TaskDTO, error) {
	t.log.Info("paginate without status", zap.Int("offset", offset), zap.Int("limit", limit), zap.String("status1", ignored.String()))

	const query = `SELECT id, value, created_at, status
FROM tasks
WHERE status <> $1
ORDER BY created_at DESC
OFFSET $2 LIMIT $3;`

	return t.getManyTasksByQuery(ctx, query, ignored.Int(), offset, limit)
}

func (t *TaskRepository) PaginateFilterWithStatus(ctx context.Context, offset int, limit int, task string, status model.Status) ([]model.TaskDTO, error) {
	t.log.Info("paginate without statuses", zap.Int("offset", offset), zap.Int("limit", limit), zap.String("status1", status.String()))

	const query = `SELECT id, value, created_at, status
FROM tasks
WHERE value LIKE CONCAT('%%',$1::text,'%%') AND status = $2
ORDER BY created_at DESC
OFFSET $3 LIMIT $4;`

	return t.getManyTasksByQuery(ctx, query, task, status.Int(), offset, limit)
}

func (t *TaskRepository) PaginateFilterWithStatuses(ctx context.Context, offset int, limit int, task string, status1, status2 model.Status) ([]model.TaskDTO, error) {
	const query = `SELECT id, value, created_at, status
FROM tasks
WHERE value LIKE CONCAT('%%',$1::text,'%%') AND (status = $2 OR status = $3)
ORDER BY created_at DESC
OFFSET $4 LIMIT $5;`
	return t.getManyTasksByQuery(ctx, query, task, status1.Int(), status2.Int(), offset, limit)
}

func (t *TaskRepository) PaginateFilterWithoutStatus(ctx context.Context, offset int, limit int, task string, ignored model.Status) ([]model.TaskDTO, error) {
	const query = `SELECT id, value, created_at, status
FROM tasks
WHERE value LIKE CONCAT('%%',$1::text,'%%') AND status <> $2
ORDER BY created_at DESC
OFFSET $3 LIMIT $4;`
	return t.getManyTasksByQuery(ctx, query, task, ignored.Int(), offset, limit)
}
