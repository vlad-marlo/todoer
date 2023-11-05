package memory

import (
	"sync"

	"go.uber.org/zap"
)

// Option is implementation option func pattern in storage configurator.
type Option interface {
	apply(storage *Storage)
}

// FuncOption is adapter to funcion to be
// valid Option.
type FuncOption func(*Storage)

// apply implements Option interface.
func (opt FuncOption) apply(storage *Storage) {
	storage.mu.Lock()
	opt(storage)
	storage.mu.Unlock()
}

func applyAsync(wg *sync.WaitGroup, storage *Storage, option Option) {
	defer wg.Done()
	option.apply(storage)
}

// WithZapLogger adds logger to storage if not provided.
func WithZapLogger(log *zap.Logger) Option {
	return FuncOption(func(s *Storage) {
		s.log = log
		s.tasks.log = log
	})
}
