package postgresql

import (
	"context"

	"github.com/cyborgvova/training/ozon/example/internal/repository"
	"github.com/cyborgvova/training/ozon/example/pkg/db"
)

type UsersRepo struct {
	db db.DBops
}

func NewUsersRepo(db db.DBops) *UsersRepo {
	return &UsersRepo{db: db}
}

func (u *UsersRepo) Add(ctx context.Context, user *repository.User) (int64, error) {
	var id int64
	err := u.db.ExecQueryRow(ctx, "INSERT INTO users(name) VALUES ($1) RETURNING id", user.Name).Scan(&id)
	return id, err
}

func (u *UsersRepo) GetById(ctx context.Context, id int64) (*repository.User, error) {
	var user *repository.User
	err := u.db.Get(ctx, user, "SELECT id, name, created_at, update_at FROM users WHERE id=$1", id)
	if err != nil {
		return nil, repository.ErrObjectNotFound
	}
	return user, err
}

func (u *UsersRepo) List(ctx context.Context) ([]*repository.User, error) {
	users := make([]*repository.User, 0)
	err := u.db.Select(ctx, &users, "SELECT id, name, created_at, updated_at FROM users")
	if err != nil {
		return nil, err
	}
	return users, nil
}

func (u *UsersRepo) Update(ctx context.Context, user *repository.User) (bool, error) {
	conn, err := u.db.Exec(ctx, "UPDATE users SET (name, updated_at) = ($1, now()) WHERE id=$2 RETURNING id", user.Name, user.Id)
	return conn.Update(), err
}
