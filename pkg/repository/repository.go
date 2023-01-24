package repository

import (
	"github.com/dahaev/todo.git"
	"github.com/jmoiron/sqlx"
)

type Authorization interface {
	CreateUser(user todo.User) (int, error)
	GetUser(user, password string) (todo.User, error)
}

type TodoList interface {
	Create(userId int, list todo.Todolist) (int, error)
}

type TodoItem interface {
}

type Repository struct {
	Authorization
	TodoItem
	TodoList
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: NewAuthPostgres(db),
	}
}
