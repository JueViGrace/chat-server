package server

import (
	"context"
	"gps-tracker/internal/api"
	"gps-tracker/internal/data"
)

type Server interface {
	Init() error
	ShutDown(ctx context.Context) error
}

type server struct {
	api api.Api
	db  data.Storage
}

func New() Server {
	server := &server{
		api: api.New(),
		db:  data.NewStorage(),
	}
	return server
}

func (s *server) Init() error {
	err := s.api.Start()
	if err != nil {
		dbErr := s.db.Close()
		if dbErr != nil {
			return err
		}
		return err
	}
	return nil
}

func (s *server) ShutDown(ctx context.Context) error {
	err := s.api.ShutDownWithContext(ctx)
	if err != nil {
		return err
	}
	err = s.db.Close()
	if err != nil {
		return err
	}
	return nil
}
