package service

import (
	"context"
	"github.com/vlad-marlo/todoer/internal/model"
)

func (s *Service) Create(ctx context.Context, task string, status model.Status) (*model.CreateTaskResponse, error) {
	//TODO implement me
	panic("implement me")
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

func (s *Service) ChangeStatus(ctx context.Context, id, status string) (*model.TaskDTO, error) {
	//TODO implement me
	panic("implement me")
}

func (s *Service) CreateMany(ctx context.Context, tasks []model.TaskDTO) error {
	//TODO implement me
	panic("implement me")
}
