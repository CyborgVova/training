package main

import (
	"context"
	"fmt"
	"log"

	"github.com/cyborgvova/training/ozon/example/internal/repository"
	"github.com/cyborgvova/training/ozon/example/internal/repository/inmemory"
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
	fmt.Printf("added new record:\nid: %d user: '%s'\n", id, user.Name)

	userRepo2 := inmemory.NewUsersRepo()
	id2, err2 := userRepo2.Add(ctx, &user)
	userRepo2.Add(ctx, &repository.User{Name: "July"})
	userRepo2.Add(ctx, &repository.User{Name: "Vova"})
	if err != nil {
		log.Fatalf("error add user: %v", err2)
	}
	fmt.Printf("added new record:\nid: %d user: '%s'\n", id2, user.Name)

	user2 := repository.User{
		Id:   1,
		Name: "Semion",
	}

	ok, err := userRepo.Update(ctx, &user2)
	if err != nil {
		log.Fatalf("error update user: %v", err)
	}

	fmt.Printf("update record: %v\nid: %d user: '%s'\n", ok, user.Id, user.Name)

	ok2, err2 := userRepo2.Update(ctx, &user2)
	if err2 != nil {
		log.Fatalf("error update user: %v", err2)
	}

	fmt.Printf("update record: %v\nid: %d user: '%s'\n", ok2, user.Id, user.Name)

	list, _ := userRepo.List(ctx)
	list2, _ := userRepo2.List(ctx)
	for _, v := range list {
		fmt.Println(*v)
	}

	for _, v := range list2 {
		fmt.Println(*v)
	}

}
