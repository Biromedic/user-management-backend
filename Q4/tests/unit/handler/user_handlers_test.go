package handler_test

import (
	"Q4/internal/model"
	"Q4/internal/service"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

type MockUserRepository struct {
	mock.Mock
}

func (m *MockUserRepository) GetAllUsers() ([]model.User, error) {
	args := m.Called()
	return args.Get(0).([]model.User), args.Error(1)
}

func (m *MockUserRepository) GetUserByID(id int) (*model.User, error) {
	args := m.Called(id)
	return args.Get(0).(*model.User), args.Error(1)
}

func (m *MockUserRepository) CreateUser(user *model.User) error {
	args := m.Called(user)
	return args.Error(0)
}

func (m *MockUserRepository) UpdateUser(user *model.User) error {
	args := m.Called(user)
	return args.Error(0)
}

func (m *MockUserRepository) DeleteUser(id int) error {
	args := m.Called(id)
	return args.Error(0)
}

func TestUserService_GetAllUsers(t *testing.T) {
	mockRepo := new(MockUserRepository)
	mockRepo.On("GetAllUsers").Return([]model.User{
		{ID: 1, Name: "Ahmet", Email: "ahmet@example.com"},
	}, nil)

	userService := service.NewUserService(mockRepo)
	users, err := userService.GetAllUsers()

	assert.NoError(t, err)
	assert.Len(t, users, 1)
	assert.Equal(t, "Ahmet", users[0].Name)
	mockRepo.AssertExpectations(t)
}

func TestUserService_CreateUser(t *testing.T) {
	mockRepo := new(MockUserRepository)
	mockRepo.On("CreateUser", mock.Anything).Return(nil)

	userService := service.NewUserService(mockRepo)
	err := userService.CreateUser(&model.User{Name: "Ahmet", Email: "ahmet@example.com"})

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestUserService_GetUserByID(t *testing.T) {
	mockRepo := new(MockUserRepository)
	mockRepo.On("GetUserByID", 1).Return(&model.User{ID: 1, Name: "Ahmet", Email: "ahmet@example.com"}, nil)

	userService := service.NewUserService(mockRepo)
	user, err := userService.GetUserByID(1)

	assert.NoError(t, err)
	assert.Equal(t, "Ahmet", user.Name)
	mockRepo.AssertExpectations(t)
}

func TestUserService_DeleteUser(t *testing.T) {
	mockRepo := new(MockUserRepository)
	mockRepo.On("DeleteUser", 1).Return(nil)

	userService := service.NewUserService(mockRepo)
	err := userService.DeleteUser(1)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}
