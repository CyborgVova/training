package repository

import (
	"fmt"

	"github.com/jackc/pgtype"
)

type User struct {
	Id        int64              `db:"id"`
	Name      string             `db:"name"`
	CreatedAt pgtype.Timestamptz `db:"created_at"`
	UpdatedAt pgtype.Timestamptz `db:"updated_at"`
}

func (u User) String() string {
	return fmt.Sprintf("%d, %s, %v, %v", u.Id, u.Name, u.CreatedAt, u.UpdatedAt)
}
