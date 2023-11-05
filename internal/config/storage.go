package config

import (
	"context"
	"fmt"
)

type Storage struct {
	DatabaseURI      string `config:"database-uri,short=d"`
	DatabaseHost     string `config:"database-host"`
	DatabasePort     int    `config:"database-port"`
	DatabaseName     string `config:"database-name"`
	DatabaseUser     string `config:"database-user"`
	DatabasePassword string `config:"database-password"`
}

func NewStorage() (*Storage, error) {
	s := new(Storage)
	if err := getConfitaLoader().Load(context.Background(), s); err != nil {
		return nil, fmt.Errorf("confita: unable to load storage config: %w", err)
	}
	if s.DatabaseURI == "" {
		s.DatabaseURI = fmt.Sprintf(
			"postgresql://%s:%s@%s:%d/%s?sslmode=disable",
			s.DatabaseUser, s.DatabasePassword, s.DatabaseHost, s.DatabasePort, s.DatabaseName,
		)

	}
	return s, nil
}
