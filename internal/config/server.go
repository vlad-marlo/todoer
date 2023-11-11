package config

import "context"

type Server struct {
	BindAddr string `config:"bind-addr"`
	BindPort int    `config:"bind-port,short=p"`
}

func NewServer() (*Server, error) {
	cfg := &Server{
		BindAddr: "localhost",
		BindPort: 8080,
	}
	if err := getConfitaLoader().Load(context.Background(), cfg); err != nil {
		return nil, err
	}
	return cfg, nil
}
