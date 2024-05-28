package db

import (
	"context"
	"fmt"
	"log"

	"github.com/jackc/pgx/v4/pgxpool"
)

const (
	driver   = "postgres"
	host     = "localhost"
	port     = "5432"
	user     = "test"
	password = "test"
	dbname   = "test"
)

func NewDB(ctx context.Context) (*Database, error) {
	dsn := fmt.Sprintf("%s://%s:%s@%s:%s/%s", driver, user, password, host, port, dbname)
	fmt.Println(dsn)
	cluster, err := pgxpool.Connect(ctx, dsn)
	if err != nil {
		log.Fatal(err)
	}
	return newDatabase(cluster), nil
}
