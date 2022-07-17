package service

import (
	todo "github.com/Vladosya/go-test-rest"
	"github.com/Vladosya/go-test-rest/pkg/repository"
)

type PostService struct {
	repo repository.TodoPost
}

func NewPostService(r repository.TodoPost) *PostService {
	return &PostService{
		repo: r,
	}
}

func (s *PostService) CreatePost(p todo.Post) (int, error) {
	return s.repo.CreatePost(p)
}

func (s *PostService) GetPosts() ([]todo.Post, error) {
	return s.repo.GetPosts()
}

func (s *PostService) DeletePostById(id string) ([]todo.Post, error) {
	return s.repo.DeletePostById(id)
}
