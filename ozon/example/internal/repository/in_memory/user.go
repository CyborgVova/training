package in_memory

import (
	"github.com/cyborgvova/training/ozon/example/internal/repository"
)

type UsersRepo struct {
	users map[int64]repository.User
}

func NewUsersRepo() *UsersRepo {
	return &UsersRepo{users: make(map[int64]repository.User)}
}
