package database

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

type UserRepository interface {
	GetByID(table string, id int) []byte
}

type Postgres struct {
	DB *sql.DB
}

func New() *sql.DB {
	db, err := sql.Open("postgres", "postgres://user:password@localhost:5432/dbname")
	if err != nil {
		log.Fatalf("error create to database connection: %v", err)
	}
	return db
}

func NewUserRepository(db *sql.DB) UserRepository {
	return &Postgres{DB: db}
}

func (p *Postgres) GetByID(table string, id int) []byte {
	res, err := p.DB.Query("select * from $1 where id = $2", table, id)
	if err != nil {
		log.Printf("error get by id:", err)
	}

	for res.Next() {
		res.Scan()
	}
}
