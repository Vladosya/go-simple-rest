package service

import (
	todo "github.com/Vladosya/go-test-rest"
	"github.com/Vladosya/go-test-rest/pkg/repository"
)

type UserService struct {
	repo repository.TodoUser
}

func NewUserService(r repository.TodoUser) *UserService {
	return &UserService{
		repo: r,
	}
}

func (s *UserService) CreateUser(u todo.User) (int, error) {
	return s.repo.CreateUser(u)
}

func (s *UserService) GetUsers() ([]todo.User, error) {
	return s.repo.GetUsers()
}

func (s *UserService) GetUserById(id string) ([]todo.User, error) {
	return s.repo.GetUserById(id)
}

func (s *UserService) DeleteUserById(id string) ([]todo.User, error) {
	return s.repo.DeleteUserById(id)
}
