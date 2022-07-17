package repository

import (
	todo "github.com/Vladosya/go-test-rest"
	"github.com/jmoiron/sqlx"
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

type Repository struct {
	TodoUser
	TodoPost
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		TodoUser: NewUserPostgres(db),
		TodoPost: NewPostPostgres(db),
	}
}
