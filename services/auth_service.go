package services

import (
	"fmt"

	"github.com/guiidc/todo-list/entities"
)

type AuthService interface {
	Register(newUser entities.CreateUserRequest) bool
}

type authService struct {
}

func NewAuthService() AuthService {
	return &authService{}
}

func (as *authService) Register(newUser entities.CreateUserRequest) bool {
	fmt.Println("REGISTER")
	return true
}
