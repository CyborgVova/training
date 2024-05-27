package usecase

import (
	"github.com/cyborgvova/training/cleanarchitecture/pkg/database"
)

type UserUseCase struct{}

func NewUserUseCase(userRepo *database.UserRepository) *UserUseCase {
	return &UserUseCase{}
}
