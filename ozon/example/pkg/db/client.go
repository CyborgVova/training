package db

import (
	"context"
	"log"

	"github.com/jackc/pgx/v4/pgxpool"
)

const (
	host     = "localhost"
	port     = "5432"
	user     = "test"
	password = "test"
	dbname   = "test"
)

func NewDB(ctx context.Context) (*Database, error) {
	dsn := "host=localhost port=5432 dbname=test user=test password=test sslmode=disable"
	cluster, err := pgxpool.Connect(ctx, dsn)
	if err != nil {
		log.Fatal(err)
	}
	return newDatabase(cluster), nil
}
