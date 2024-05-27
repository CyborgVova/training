package main

import (
	"fmt"

	"github.com/cyborgvova/training/cleanarchitecture/internal/delivery/http"
	"github.com/cyborgvova/training/cleanarchitecture/internal/usecase"
	"github.com/cyborgvova/training/cleanarchitecture/pkg/database"
)

func main() {
	db := database.New()
	userRepo := database.NewUserRepository(db)
	userUC := usecase.NewUserUseCase(userRepo)
	userHandler := http.NewUserHandler(userUC)

	fmt.Println("Server started at :8080")
	userHandler.Run(":8080")
}
