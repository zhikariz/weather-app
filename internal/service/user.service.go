package service

import (
	"context"

	"github.com/zhikariz/weather-app/entity"
)

type UserUseCase interface {
	FindAll(ctx context.Context) ([]*entity.User, error)
}

type UserRepository interface {
	FindAll(ctx context.Context) ([]*entity.User, error)
}

type UserService struct {
	repository UserRepository
}

func NewUserService(repository UserRepository) *UserService {
	return &UserService{repository}
}

func (s *UserService) FindAll(ctx context.Context) ([]*entity.User, error) {
	return s.repository.FindAll(ctx)
}
