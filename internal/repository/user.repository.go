package repository

import (
	"context"

	"github.com/zhikariz/weather-app/entity"
	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{
		db: db,
	}
}

func (r *UserRepository) FindAll(ctx context.Context) ([]*entity.User, error) {
	users := make([]*entity.User, 0)
	err := r.db.WithContext(ctx).Find(&users).Error // SELECT * FROM users
	if err != nil {
		return nil, err
	}
	return users, nil
}
