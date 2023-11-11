//go:generate easyjson -all task.go
package model

import (
	"github.com/google/uuid"
	"time"
)

type TaskDTO struct {
	ID        uuid.UUID `json:"id"`
	Value     string    `json:"value"`
	CreatedAt time.Time `json:"created_at"`
	Status    Status    `json:"status"`
}
