package service

import (
	"errors"

	"user-service/model"
	"user-service/repository"
)

type UserService interface {
	List(page, size int) []model.User
	Create(name string) (model.User, error)
	GetByID(id int) (model.User, error)
}

type userService struct {
	repo repository.UserRepository
}

func NewUserService(repo repository.UserRepository) UserService {
	return &userService{repo: repo}
}

func (s *userService) List(page, size int) []model.User {
	users := s.repo.List(page, size)
	return users
}

func (s *userService) Create(name string) (model.User, error) {
	if name == "" {
		return model.User{}, errors.New("name is required")
	}
	newUser, err := s.repo.Create(name)
	if err != nil {
		return model.User{}, err
	}
	return newUser, nil
}

func (s *userService) GetByID(id int) (model.User, error) {
	user, err := s.repo.GetByID(id)
	if err != nil {
		return model.User{}, err
	}
	return user, nil
}
