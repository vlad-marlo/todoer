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

func (s *Service) prepareGetTasksResponse(ctx context.Context, offset, limit int) (*model.GetTasksResponse, error) {
	count, err := s.storage.Task().Count(ctx)
	if err != nil {
		return nil, &model.ErrorMessage{
			Endpoint: "/api/v1/tasks",
			Code:     http.StatusInternalServerError,
			Status:   fmt.Sprintf("unable to get count of stored tasks: %v", err),
		}
	}

	resp := model.GetTasksResponse{
		Count: count,
		Next:  fmt.Sprintf("/api/v1/tasks?offset=%d&limit=%d", offset+limit, limit),
		Tasks: nil,
	}
	if offset < limit {
		resp.Previous = fmt.Sprintf("/api/v1/tasks?offset=0&limit=%d", offset)
	} else {
		resp.Previous = fmt.Sprintf("/api/v1/tasks?offset=%d&limit=%d", offset-limit, limit)
	}
	return &resp, nil
}

func (s *Service) getManyNoStatus(ctx context.Context, offset, limit int, task string) (*model.GetTasksResponse, error) {
	resp, err := s.prepareGetTasksResponse(ctx, offset, limit)
	if err != nil {
		return nil, err
	}

	if task != "" {
		resp.Tasks, err = s.storage.Task().PaginateFilter(ctx, offset, limit, task)
	} else {
		resp.Tasks, err = s.storage.Task().Paginate(ctx, offset, limit)
	}

	if err != nil {
		return nil, &model.ErrorMessage{
			Endpoint: "/api/v1/tasks GET",
			Code:     http.StatusBadRequest,
			Status:   fmt.Sprintf("some error in storage layer: %v", err),
		}
	}

	return resp, nil
}

func (s *Service) getManyOneStatus(ctx context.Context, offset, limit int, task string, status model.Status) (*model.GetTasksResponse, error) {
	resp, err := s.prepareGetTasksResponse(ctx, offset, limit)
	if err != nil {
		return nil, err
	}

	if task != "" {
		resp.Tasks, err = s.storage.Task().PaginateFilterWithStatus(ctx, offset, limit, task, status)
	} else {
		resp.Tasks, err = s.storage.Task().PaginateWithStatus(ctx, offset, limit, status)
	}

	if err != nil {
		return nil, &model.ErrorMessage{
			Endpoint: "/api/v1/tasks GET",
			Code:     http.StatusInternalServerError,
			Status:   fmt.Sprintf("got unexpected error: %v", err),
		}
	}

	return resp, nil
}

func (s *Service) getManyTwoStatus(ctx context.Context, offset, limit int, task string, status1, status2 model.Status) (*model.GetTasksResponse, error) {
	resp, err := s.prepareGetTasksResponse(ctx, offset, limit)
	if err != nil {
		return nil, err
	}

	if task != "" {
		resp.Tasks, err = s.storage.Task().PaginateFilterWithStatuses(ctx, offset, limit, task, status1, status2)
	} else {
		resp.Tasks, err = s.storage.Task().PaginateWithStatuses(ctx, offset, limit, status1, status2)
	}

	if err != nil {
		return nil, &model.ErrorMessage{
			Endpoint: "/api/v1/tasks GET",
			Code:     http.StatusInternalServerError,
			Status:   fmt.Sprintf("got unexpected error: %v", err),
		}
	}

	return resp, nil
}

func (s *Service) getManyThreeStatuses(ctx context.Context, offset, limit int, task string, statuses ...model.Status) (*model.GetTasksResponse, error) {
	resp, err := s.prepareGetTasksResponse(ctx, offset, limit)
	if err != nil {
		return nil, err
	}
	var ignored model.Status
	ignored.SetStatusWhichDoesNotContainsInStatuses(statuses...)

	if task != "" {
		resp.Tasks, err = s.storage.Task().PaginateFilterWithoutStatus(ctx, offset, limit, task, ignored)
	} else {
		resp.Tasks, err = s.storage.Task().PaginateWithoutStatus(ctx, offset, limit, ignored)
	}

	if err != nil {
		return nil, &model.ErrorMessage{
			Endpoint: "/api/v1/tasks GET",
			Code:     http.StatusInternalServerError,
			Status:   fmt.Sprintf("got unexpected error: %v", err),
		}
	}

	return resp, nil
}

func (s *Service) GetMany(ctx context.Context, offset, limit int, task string, statuses ...model.Status) (*model.GetTasksResponse, error) {
	switch len(statuses) {
	case 0, 4:
		return s.getManyNoStatus(ctx, offset, limit, task)
	case 1:
		return s.getManyOneStatus(ctx, offset, limit, task, statuses[0])
	case 2:
		return s.getManyTwoStatus(ctx, offset, limit, task, statuses[0], statuses[1])
	case 3:
		return s.getManyThreeStatuses(ctx, offset, limit, task, statuses...)
	default:
		return nil, model.ErrorMessage{
			Endpoint: "/api/v1/tasks GET",
			Code:     http.StatusBadRequest,
			Status:   "unacceptable amount of statuses",
		}
	}
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
