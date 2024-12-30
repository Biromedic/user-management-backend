package service

import (
	"Q4/internal/model"
	"Q4/internal/repository"
)

type UserServiceInterface interface {
	GetAllUsers() ([]model.User, error)
	GetUserByID(id int) (*model.User, error)
	CreateUser(user *model.User) error
	UpdateUser(user *model.User) error
	DeleteUser(id int) error
}

type UserService struct {
	Repo repository.UserRepository
}

func NewUserService(repo repository.UserRepository) UserServiceInterface {
	return &UserService{
		Repo: repo,
	}
}

func (s *UserService) GetAllUsers() ([]model.User, error) {
	return s.Repo.GetAllUsers()
}

func (s *UserService) GetUserByID(id int) (*model.User, error) {
	return s.Repo.GetUserByID(id)
}

func (s *UserService) CreateUser(user *model.User) error {
	return s.Repo.CreateUser(user)
}

func (s *UserService) UpdateUser(user *model.User) error {
	return s.Repo.UpdateUser(user)
}

func (s *UserService) DeleteUser(id int) error {
	return s.Repo.DeleteUser(id)
}
