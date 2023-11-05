package memory

import (
	"sync"

	"github.com/google/uuid"
	"github.com/vlad-marlo/todoer/internal/model"
	"go.uber.org/zap"
)

type repository struct {
	mu   sync.Mutex
	log  *zap.Logger
	data map[uuid.UUID]*model.TaskDTO
}
