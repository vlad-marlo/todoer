package service

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"github.com/vlad-marlo/todoer/internal/model"
	"net/http"
	"time"
)

func (s *Service) Create(ctx context.Context, taskValue string, status model.Status) (*model.CreateTaskResponse, error) {
	var (
		now  = time.Now().UTC()
		id   = uuid.New()
		task = &model.TaskDTO{
			ID:        id,
			Value:     taskValue,
			CreatedAt: now,
			Status:    status,
		}
	)

	err := s.storage.Task().Create(ctx, task)
	if err != nil {
		return nil, model.ErrorMessage{
			Endpoint: "/api/v1/tasks POST",
			Code:     http.StatusInternalServerError,
			Status:   fmt.Sprintf("unknown error in storage layer: %v", err),
		}
	}

	return &model.CreateTaskResponse{
		ID:        id,
		Task:      taskValue,
		CreatedAt: now,
	}, nil
}

func (s *Service) GetMany(ctx context.Context, offset, limit int, task string, statuses ...model.Status) (*model.GetTasksResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (s *Service) GetOne(ctx context.Context, id string) (*model.TaskDTO, error) {
	//TODO implement me
	panic("implement me")
}

func (s *Service) Update(ctx context.Context, id string, task string, status model.Status) (*model.TaskDTO, error) {
	//TODO implement me
	panic("implement me")
}

func (s *Service) Delete(ctx context.Context, id string) error {
	//TODO implement me
	panic("implement me")
}

func (s *Service) ChangeStatus(ctx context.Context, id string, status model.Status) (*model.TaskDTO, error) {
	//TODO implement me
	panic("implement me")
}

func (s *Service) CreateMany(ctx context.Context, tasks []model.TaskDTO) error {
	//TODO implement me
	panic("implement me")
}
