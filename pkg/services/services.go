package services

import (
	"github.com/dahaev/todo.git"
	"github.com/dahaev/todo.git/pkg/repository"
)

// pkg/services/go generate
//
//go:generate mockgen -source=services.go -destination=mocks/mock.go
type Authorization interface {
	CreateUser(user todo.User) (int, error)
	GenerateToken(username, password string) (string, error)
	ParseToken(token string) (int, error)
}

type TodoList interface {
	Create(userId int, list todo.Todolist) (int, error)
}

type TodoItem interface {
}

type Services struct {
	Authorization
	TodoItem
	TodoList
}

func NewService(repo *repository.Repository) *Services {
	return &Services{
		Authorization: NewAuthService(repo.Authorization),
		TodoList:      NewTodoListService(repo.TodoList),
	}
}
