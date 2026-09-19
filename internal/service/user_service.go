package service

import (
	"context"
	"day01/internal/model"
	"day01/internal/store"
)

type UserService struct{ store *store.UserStore }

func NewUserService(s *store.UserStore) *UserService                       { return &UserService{store: s} }
func (s *UserService) ListUsers(ctx context.Context) ([]model.User, error) { return s.store.List(ctx) }
func (s *UserService) CreateUser(ctx context.Context, req model.CreateUserRequest) (*model.User, error) {
	return s.store.Create(ctx, req.Username, req.Email)
}
