package service

import (
	"context"
	"errors"

	"github.com/zhikariz/weather-app/entity"
)

type UserUseCase interface {
	FindAll(ctx context.Context) ([]*entity.User, error)
	Create(ctx context.Context, user *entity.User) error
	Update(ctx context.Context, user *entity.User) error
	Delete(ctx context.Context, id int64) error
	FindByID(ctx context.Context, id int64) (*entity.User, error)
}

type UserRepository interface {
	FindAll(ctx context.Context) ([]*entity.User, error)
	Create(ctx context.Context, user *entity.User) error
	Update(ctx context.Context, user *entity.User) error
	Delete(ctx context.Context, id int64) error
	FindByID(ctx context.Context, id int64) (*entity.User, error)
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

func (s *UserService) Create(ctx context.Context, user *entity.User) error {
	return s.repository.Create(ctx, user)
}

func (s *UserService) Update(ctx context.Context, user *entity.User) error {
	if user.Role != "" {
		if user.Role != "Administrator" && user.Role != "Editor" {
			return errors.New("role harus di isi Administrator / Editor")
		}
	}
	return s.repository.Update(ctx, user)
}

func (s *UserService) Delete(ctx context.Context, id int64) error {
	return s.repository.Delete(ctx, id)
}

func (s *UserService) FindByID(ctx context.Context, id int64) (*entity.User, error) {
	return s.repository.FindByID(ctx, id)
}
