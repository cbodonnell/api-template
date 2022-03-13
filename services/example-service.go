package services

import (
	"github.com/cbodonnell/api-template/config"
	"github.com/cbodonnell/api-template/models"
	"github.com/cbodonnell/api-template/repositories"
	"github.com/cbodonnell/api-template/workers"
)

type ExampleService struct {
	conf   config.Configuration
	repo   repositories.Repository
	worker workers.Worker
}

func NewExampleService(_conf config.Configuration, _repo repositories.Repository) Service {
	return &ExampleService{
		conf: _conf,
		repo: _repo,
	}
}

func (s *ExampleService) GetUserByName(name string) (models.User, error) {
	return s.repo.QueryUserByName(name)
}
