package inmemory

import (
	"context"
	"time"

	"github.com/cyborgvova/training/ozon/example/internal/repository"
)

var count int64

type UsersRepo struct {
	users map[int64]*repository.User
}

func NewUsersRepo() *UsersRepo {
	return &UsersRepo{users: make(map[int64]*repository.User)}
}

func (u *UsersRepo) Add(ctx context.Context, user *repository.User) (int64, error) {
	count++
	user.CreatedAt.Time = time.Now().UTC().Local()
	user.Id = count
	u.users[user.Id] = user
	return user.Id, nil
}
func (u *UsersRepo) GetById(ctx context.Context, id int64) (*repository.User, error) {
	val, ok := u.users[id]
	if !ok {
		return nil, repository.ErrObjectNotFound
	}
	return val, nil
}
func (u *UsersRepo) List(ctx context.Context) ([]*repository.User, error) {
	if len(u.users) == 0 {
		return nil, repository.ErrObjectNotFound
	}

	users := make([]*repository.User, 0)
	for _, v := range u.users {
		users = append(users, v)
	}

	return users, nil
}
func (u *UsersRepo) Update(ctx context.Context, user *repository.User) (bool, error) {
	user.UpdatedAt.Time = time.Now().UTC().Local()
	if _, ok := u.users[user.Id]; !ok {
		return false, repository.ErrObjectNotFound
	}

	user.CreatedAt = u.users[user.Id].CreatedAt
	u.users[user.Id] = user

	return true, nil
}
