package service

import (
	todo "github.com/KDias-code/todoapp"
	"github.com/KDias-code/todoapp/pkg/repository"
)

type Autharization interface {
	CreateUser(user todo.User) (int, error)
}

type TodoList interface {
}

type TodoItem interface {
}

type Service struct {
	Autharization
	TodoList
	TodoItem
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Autharization: NewAuthService(repos.Autharization),
	}
}