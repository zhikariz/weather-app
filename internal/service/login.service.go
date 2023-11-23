package service

import (
	"context"
	"errors"

	"github.com/zhikariz/weather-app/entity"
)

type LoginUseCase interface {
	Login(ctx context.Context, email, password string) (*entity.User, error)
}

type LoginRepository interface {
	FindByEmail(ctx context.Context, email string) (*entity.User, error)
}

type LoginService struct {
	repo LoginRepository
}

func NewLoginService(repo LoginRepository) *LoginService {
	return &LoginService{
		repo: repo,
	}
}

func (s *LoginService) Login(ctx context.Context, email, password string) (*entity.User, error) {
	user, err := s.repo.FindByEmail(ctx, email)

	if err != nil {
		return nil, err
	}

	if user == nil {
		return nil, errors.New("user with that email not found")
	}

	if user.Password != password {
		return nil, errors.New("incorrect login credentials")
	}

	return user, nil
}
