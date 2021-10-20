package service

import (
	"errors"
	"main/entity"
	"main/repository"
	"math/rand"
)

type PostService interface {
	Validate(post *entity.Post) error
	Create(post *entity.Post) (*entity.Post, error)
	FindAll() ([]entity.Post, error)
}

type service struct{}

var (
	repo repository.PostRepository
)

func NewPostService(r repository.PostRepository) PostService {
	repo = r
	return &service{}
}

func (*service) Validate(post *entity.Post) error {
	if post == nil {
		err := errors.New("the post is nil")
		return err
	}
	if post.Title == "" {
		err := errors.New("the post title is empty")
		return err
	}
	return nil

}
func (*service) Create(post *entity.Post) (*entity.Post, error) {
	post.Id = rand.Int()
	return repo.Save(post)

}
func (*service) FindAll() ([]entity.Post, error) {
	return repo.FindAll()
}
