package repository

import (
	"context"
	"errors"
)

var (
	ErrObjectNotFound = errors.New("object not found")
)

type UserRepo interface {
	Add(ctx context.Context, user *User) (int64, error)
	GetById(ctx context.Context) (*User, error)
	List(ctx context.Context)
	Update(ctx context.Context)
}
