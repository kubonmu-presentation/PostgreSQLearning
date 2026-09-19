package store

import (
	"context"
	"day01/internal/model"
	"github.com/jackc/pgx/v5/pgxpool"
)

type UserStore struct{ db *pgxpool.Pool }

func NewUserStore(db *pgxpool.Pool) *UserStore { return &UserStore{db: db} }
func (s *UserStore) List(ctx context.Context) ([]model.User, error) {
	rows, err := s.db.Query(ctx, `SELECT id, username, email, create_at FROM users ORDER BY id`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	users := make([]model.User, 0)
	for rows.Next() {
		var u model.User
		if err := rows.Scan(&u.ID, &u.Username, &u.Email, &u.CreateAt); err != nil {
			return nil, err
		}
		users = append(users, u)
	}
	return users, rows.Err()
}
func (s *UserStore) Create(ctx context.Context, username, email string) (*model.User, error) {
	var u model.User
	err := s.db.QueryRow(ctx, `INSERT INTO public.users (username, email) VALUES ($1, $2) RETURNING id, username, email, create_at`, username, email).Scan(&u.ID, &u.Username, &u.Email, &u.CreateAt)
	if err != nil {
		return nil, err
	}
	return &u, nil
}
