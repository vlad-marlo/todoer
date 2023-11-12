package config

import (
	"context"
	"fmt"
	"os"
	"strings"
)

type Storage struct {
	DatabaseURI          string `config:"database-uri,short=d"`
	DatabaseHost         string `config:"database-host"`
	DatabasePort         int    `config:"database-port"`
	DatabaseName         string `config:"database-name"`
	DatabaseUser         string `config:"database-user"`
	DatabasePassword     string `config:"database-password"`
	DatabasePasswordFile string `config:"database-password-file,short=f"`
}

func (cfg *Storage) getPassword() error {
	pass, err := os.ReadFile(cfg.DatabasePasswordFile)
	if err != nil {
		return fmt.Errorf("read data from file: %w", err)
	}
	cfg.DatabasePassword = strings.Trim(string(pass), " \n")
	return nil
}

//goland:noinspection ALL
func (cfg *Storage) load() error {

	if err := getConfitaLoader().Load(context.Background(), cfg); err != nil {
		return fmt.Errorf("confita: unable to load storage config: %w", err)
	}

	if cfg.DatabasePasswordFile != "" {
		if err := cfg.getPassword(); err != nil {
			return fmt.Errorf("get password from file: %w", err)
		}
	}

	cfg.setURI()

	return nil
}

func (cfg *Storage) setURI() {
	if cfg.DatabaseURI == "" {
		cfg.DatabaseURI = fmt.Sprintf(
			"postgresql://%s:%s@%s:%d/%s?sslmode=disable",
			cfg.DatabaseUser, cfg.DatabasePassword, cfg.DatabaseHost, cfg.DatabasePort, cfg.DatabaseName,
		)
	}
}

//goland:noinspection ALL
func NewStorage() (*Storage, error) {
	s := &Storage{
		DatabaseHost:     "localhost",
		DatabasePort:     5432,
		DatabaseName:     "postgres",
		DatabaseUser:     "postgres",
		DatabasePassword: "postgres",
	}

	if err := s.load(); err != nil {
		return nil, fmt.Errorf("load: %w", err)
	}

	return s, nil
}

func (cfg *Storage) URI() string {
	return cfg.DatabaseURI
}
