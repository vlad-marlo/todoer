//go:generate easyjson -all responses.go

package model

import (
	"github.com/google/uuid"
	"time"
)

type (
	CreateTaskResponse struct {
		ID        uuid.UUID `json:"id"`
		Task      string    `json:"task"`
		CreatedAt time.Time `json:""`
	}
	GetTasksResponse struct {
		Count    int       `json:"count"`
		Next     string    `json:"next"`
		Previous string    `json:"previous"`
		Tasks    []TaskDTO `json:"result"`
	}
)
