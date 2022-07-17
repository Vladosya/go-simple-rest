package service

import (
	todo "github.com/Vladosya/go-test-rest"
	"github.com/Vladosya/go-test-rest/pkg/repository"
)

type TodoUser interface {
	CreateUser(u todo.User) (int, error)
	GetUsers() ([]todo.User, error)
	GetUserById(id string) ([]todo.User, error)
	DeleteUserById(id string) ([]todo.User, error)
}

type TodoPost interface {
	CreatePost(u todo.Post) (int, error)
	GetPosts() ([]todo.Post, error)
	DeletePostById(id string) ([]todo.Post, error)
}

type Service struct {
	TodoUser
	TodoPost
}

func NewService(r *repository.Repository) *Service {
	return &Service{
		TodoUser: NewUserService(r.TodoUser),
		TodoPost: NewPostService(r.TodoPost),
	}
}
