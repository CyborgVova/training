package main

import (
	"context"
	"fmt"
	"log"

	"github.com/cyborgvova/training/ozon/example/internal/repository"
	"github.com/cyborgvova/training/ozon/example/internal/repository/postgresql"
	"github.com/cyborgvova/training/ozon/example/pkg/db"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	database, err := db.NewDB(ctx)
	if err != nil {
		return
	}
	defer database.GetPool(ctx).Close()

	user := repository.User{
		Name: "poopOk",
	}

	userRepo := postgresql.NewUsersRepo(database)
	id, err := userRepo.Add(ctx, &user)
	if err != nil {
		log.Fatalf("error add user: %v", err)
	}
	fmt.Printf("added new record:\nid: %d user: '%s'", id, user.Name)

	user = repository.User{
		Id:   1,
		Name: "Semion",
	}

	ok, err := userRepo.Update(ctx, &user)
	if err != nil {
		log.Fatalf("error update user: %v", err)
	}

	fmt.Printf("update record: %v\nid: %d user: '%s'", ok, user.Id, user.Name)
}
