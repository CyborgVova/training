package in_memory

import (
	"context"

	"github.com/cyborgvova/training/ozon/example/internal/repository"
)

type UsersRepo struct {
	users map[int64]repository.User
}

func NewUsersRepo() *UsersRepo {
	return &UsersRepo{users: make(map[int64]repository.User)}
}

func (u *UsersRepo) Add(ctx context.Context, user *repository.User) (int64, error) {
	return 0, nil
}
func (u *UsersRepo) GetById(ctx context.Context) (*repository.User, error) {
	return nil, nil
}
func (u *UsersRepo) List(ctx context.Context)   {}
func (u *UsersRepo) Update(ctx context.Context) {}
