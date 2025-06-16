package service

import (
	"github.com/Dorkhan1/todo-app"
	"github.com/Dorkhan1/todo-app/pkg/repository"
)

type Authorization interface {
	CreateUser(user todo.User) (int, error)
}

type TodoList interface {
}

type TodoItem interface {
}
type Service struct {
	Authorization
	TodoList
	TodoItem
}

func NewService(repo *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repos.Authorization),
		TodoList:      repo,
		TodoItem:      repo,
	}
}
