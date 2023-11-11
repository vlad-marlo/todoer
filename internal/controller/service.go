package controller

import (
	"context"
	"github.com/vlad-marlo/todoer/internal/model"
)

type Service interface {
	Create(ctx context.Context, task string, status model.Status) (*model.CreateTaskResponse, error)
	GetMany(ctx context.Context, offset, limit int, task string, statuses ...model.Status) (*model.GetTasksResponse, error)
	GetOne(ctx context.Context, id string) (*model.TaskDTO, error)
	Update(ctx context.Context, id string, task string, status model.Status) (*model.TaskDTO, error)
	Delete(ctx context.Context, id string) error
	ChangeStatus(ctx context.Context, id, status string) (*model.TaskDTO, error)
	CreateMany(ctx context.Context, tasks []model.TaskDTO) error
}
