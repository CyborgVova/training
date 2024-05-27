package database

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

type UserRepository struct {
	DB *sql.DB
}

func New() *sql.DB {
	db, err := sql.Open("postgres", "postgres://user:password@localhost:5432/dbname")
	if err != nil {
		log.Fatalf("error create to database connection: %v", err)
	}
	return db
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{DB: db}
}
