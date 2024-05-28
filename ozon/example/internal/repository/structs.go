package repository

import (
	"fmt"
	"time"
)

type User struct {
	Id        int64     `db:"id"`
	Name      string    `db:"name"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
}

func (u User) String() string {
	return fmt.Sprintf("%d, %s, %v, %v", u.Id, u.Name, u.CreatedAt, u.UpdatedAt)
}
