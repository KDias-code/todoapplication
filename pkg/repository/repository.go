package repository

import (
	todo "github.com/KDias-code/todoapp"
	"github.com/jmoiron/sqlx"
)

type Autharization interface {
	CreateUser(user todo.User) (int, error)
}

type TodoList interface {
}

type TodoItem interface {
}

type Repository struct {
	Autharization
	TodoList
	TodoItem
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Autharization: NewAuthPostgres(db),
	}
}