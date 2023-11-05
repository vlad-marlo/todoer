package config

import (
	"sync"

	"github.com/heetch/confita"
	"github.com/heetch/confita/backend/env"
	"github.com/heetch/confita/backend/flags"
)

var (
	loader      *confita.Loader
	confitaMu   sync.Mutex
	confitaOnce sync.Once
)

func initConfita() {
	confitaMu.Lock()
	loader = confita.NewLoader(
		env.NewBackend(),
		flags.NewBackend(),
	)
	confitaMu.Unlock()
}

func getConfitaLoader() (l *confita.Loader) {
	// init confita loader if not initialized
	confitaOnce.Do(initConfita)

	confitaMu.Lock()
	l = loader
	confitaMu.Unlock()

	return l
}
