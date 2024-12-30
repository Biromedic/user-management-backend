package repository

import "Q4/internal/model"

// UserRepository defines the methods for user operations
type UserRepository interface {
	GetAllUsers() ([]model.User, error)
	GetUserByID(id int) (*model.User, error)
	CreateUser(user *model.User) error
	UpdateUser(user *model.User) error
	DeleteUser(id int) error
}
