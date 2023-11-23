package repository

import (
	"context"
	"errors"

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

func (r *UserRepository) Create(ctx context.Context, user *entity.User) error {
	if err := r.db.WithContext(ctx).Create(&user).Error; err != nil {
		return err
	}
	return nil
}

func (r *UserRepository) Update(ctx context.Context, user *entity.User) error {
	if err := r.db.WithContext(ctx).
		Model(&entity.User{}).
		Where("id = ?", user.ID).
		Updates(&user).Error; err != nil {
		return err
	}
	return nil
}

func (r *UserRepository) Delete(ctx context.Context, id int64) error {
	if err := r.db.WithContext(ctx).Delete(&entity.User{}, id).Error; err != nil {
		return err
	}
	return nil
}

func (r *UserRepository) FindByID(ctx context.Context, id int64) (*entity.User, error) {
	user := new(entity.User)
	if err := r.db.WithContext(ctx).First(&user, id).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func (r *UserRepository) FindByEmail(ctx context.Context, email string) (*entity.User, error) {
	user := new(entity.User)
	if err := r.db.WithContext(ctx).Where("email = ?", email).First(&user).Error; err != nil {
		return nil, errors.New("user with that email not found")
	}
	return user, nil
}
