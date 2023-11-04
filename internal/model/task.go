package model

import "time"

type TaskDTO struct {
	ID        string    `json:"id"`
	Value     string    `json:"value"`
	CreatedAt time.Time `json:"created_at"`
	IsDeleted bool      `json:"is_deleted"`
}

type TaskCreateDTO struct {
	ID    string `json:"id"`
	Value string `json:"value"`
}
