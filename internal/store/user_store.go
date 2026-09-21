package store

import (
	"context"
	"day01/internal/model"

	"gorm.io/gorm"
)

type UserStore struct{ db *gorm.DB }

func NewUserStore(db *gorm.DB) *UserStore { return &UserStore{db: db} }
func (s *UserStore) List(ctx context.Context) ([]model.User, error) {
	users := make([]model.User, 0)
	err := s.db.WithContext(ctx).Order("id").Find(&users).Error
	return users, err
}

func (s *UserStore) Create(ctx context.Context, username, email string) (*model.User, error) {
	u := &model.User{Username: username, Email: email}
	if err := s.db.WithContext(ctx).Create(u).Error; err != nil {
		return nil, err
	}
	return u, nil
}
