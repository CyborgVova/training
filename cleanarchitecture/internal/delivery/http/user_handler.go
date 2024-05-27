package http

import (
	"github.com/cyborgvova/training/cleanarchitecture/internal/usecase"
)

type UserHandler struct{}

func NewUserHandler(userUC *usecase.UserUseCase) UserHandler {
	return UserHandler{}
}

func (uh *UserHandler) Run(port string)
