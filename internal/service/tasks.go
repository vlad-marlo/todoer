package service

import (
	"context"
	"errors"
	"fmt"
	"github.com/google/uuid"
	"github.com/vlad-marlo/todoer/internal/model"
	"github.com/vlad-marlo/todoer/internal/storage"
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
	parsedID, err := uuid.Parse(id)
	if err != nil {
		return nil, &model.ErrorMessage{
			Endpoint: fmt.Sprintf("/api/v1/tasks/%s", id),
			Code:     http.StatusBadRequest,
			Status:   fmt.Sprintf("bad id format: %s: %v", id, err),
		}
	}
	task, err := s.storage.Task().Get(ctx, parsedID)
	if err != nil {
		if errors.Is(err, storage.ErrTaskDoesNotExists) {
			return nil, &model.ErrorMessage{
				Endpoint: fmt.Sprintf("/api/v1/tasks/%s", id),
				Code:     http.StatusNotFound,
				Status:   fmt.Sprintf("task with id=%s was not found", id),
			}
		}

		return nil, &model.ErrorMessage{
			Endpoint: fmt.Sprintf("/api/v1/tasks/%s", id),
			Code:     http.StatusInternalServerError,
			Status:   fmt.Sprintf("got unexpected error: %v", err),
		}
	}

	return task, nil
}

func (s *Service) Update(ctx context.Context, id string, value string, status model.Status) (*model.TaskDTO, error) {
	parsedID, err := uuid.Parse(id)
	if err != nil {
		return nil, &model.ErrorMessage{
			Endpoint: fmt.Sprintf("/api/v1/tasks/%s UPDATE", id),
			Code:     http.StatusBadRequest,
			Status:   fmt.Sprintf("bad id format: %s : %v", id, err),
		}
	}

	if err = s.storage.Task().Update(ctx, parsedID, value, status); err != nil {
		return nil, &model.ErrorMessage{
			Endpoint: fmt.Sprintf("/api/v1/tasks/%s UPDATE", id),
			Code:     http.StatusBadRequest,
			Status:   fmt.Sprintf("unexpected unknown error: %v", err),
		}
	}

	task, err := s.storage.Task().Get(ctx, parsedID)
	if err != nil {
		return nil, &model.ErrorMessage{
			Endpoint: fmt.Sprintf("/api/v1/tasks/%s UPDATE", id),
			Code:     http.StatusInternalServerError,
			Status:   fmt.Sprintf("unexpected unknown error: %v", err),
		}
	}

	return task, nil
}

func (s *Service) Delete(ctx context.Context, id string) error {
	parsedID, err := uuid.Parse(id)
	if err != nil {
		return &model.ErrorMessage{
			Endpoint: fmt.Sprintf("/api/v1/tasks/%s DELETE", id),
			Code:     http.StatusBadRequest,
			Status:   fmt.Sprintf("bad id format: %s : %v", id, err),
		}
	}

	if err = s.storage.Task().ChangeStatus(ctx, parsedID, model.StatusDeleted); err != nil {
		return &model.ErrorMessage{
			Endpoint: "/api/v1/%s DELETE",
			Code:     http.StatusBadRequest,
			Status:   "bad request",
		}
	}

	return nil
}

func (s *Service) ChangeStatus(ctx context.Context, id string, status model.Status) (*model.TaskDTO, error) {
	parsedID, err := uuid.Parse(id)
	if err != nil {
		return nil, &model.ErrorMessage{
			Endpoint: "/api/v1/%s/status POST",
			Code:     http.StatusBadRequest,
			Status:   fmt.Sprintf("bad id: %v", err),
		}
	}

	if err = s.storage.Task().ChangeStatus(ctx, parsedID, status); err != nil {
		return nil, &model.ErrorMessage{
			Endpoint: "/api/v1/%s/status POST",
			Code:     http.StatusBadRequest,
			Status:   "bad request",
		}
	}

	task, err := s.storage.Task().Get(ctx, parsedID)
	if err != nil {
		return nil, &model.ErrorMessage{
			Endpoint: fmt.Sprintf("/api/v1/%s/status POST", id),
			Code:     http.StatusInternalServerError,
			Status:   fmt.Sprintf("not found: %v", err),
		}
	}
	return task, nil
}

func (s *Service) CreateMany(ctx context.Context, tasks []model.TaskDTO) error {
	err := s.storage.Task().CreateMany(ctx, tasks)
	if err != nil {
		return &model.ErrorMessage{
			Endpoint: "/api/v1/tasks/file POST",
			Code:     http.StatusBadRequest,
			Status:   "bad request",
		}
	}
	return nil
}
