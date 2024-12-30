package handler_test

import (
	"Q4/internal/helpers"
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"net/http/httptest"
	"testing"

	"Q4/internal/handler"
	"Q4/internal/model"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockUserService struct {
	mock.Mock
}

func (m *MockUserService) GetAllUsers() ([]model.User, error) {
	args := m.Called()
	return args.Get(0).([]model.User), args.Error(1)
}

func (m *MockUserService) GetUserByID(id int) (*model.User, error) {
	args := m.Called(id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*model.User), args.Error(1)
}

func (m *MockUserService) CreateUser(user *model.User) error {
	args := m.Called(user)
	return args.Error(0)
}

func (m *MockUserService) UpdateUser(user *model.User) error {
	args := m.Called(user)
	return args.Error(0)
}

func (m *MockUserService) DeleteUser(id int) error {
	args := m.Called(id)
	return args.Error(0)
}

// Test functions remain the same...

// TestUserHandler_GetAllUsers tests the GetAllUsers handler
func TestUserHandler_GetAllUsers(t *testing.T) {
	mockService := new(MockUserService)
	mockService.On("GetAllUsers").Return([]model.User{{ID: 1, Name: "Test User", Email: "test@example.com"}}, nil)
	userHandler := handler.NewUserHandler(mockService)

	req, err := http.NewRequest("GET", "/users", nil)
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	userHandler.GetAllUsers(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)
	var users []model.User
	err = json.Unmarshal(rr.Body.Bytes(), &users)
	if err != nil {
		t.Fatal(err)
	}
	assert.Len(t, users, 1)
	assert.Equal(t, "Test User", users[0].Name)
}

// TestUserHandler_GetUserByID tests the GetUserByID handler with valid ID
func TestUserHandler_GetUserByID_ValidID(t *testing.T) {
	mockService := new(MockUserService)
	mockService.On("GetUserByID", 1).Return(&model.User{ID: 1, Name: "Test User", Email: "test@example.com"}, nil)
	userHandler := handler.NewUserHandler(mockService)

	router := mux.NewRouter()
	router.HandleFunc("/users/{id}", userHandler.GetUserByID).Methods("GET")

	req, err := http.NewRequest("GET", "/users/1", nil)
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	router.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)
	var user model.User
	err = json.Unmarshal(rr.Body.Bytes(), &user)
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, "Test User", user.Name)
	mockService.AssertExpectations(t)
}

// TestUserHandler_GetUserByID_NotFound tests the GetUserByID handler with invalid ID
func TestUserHandler_GetUserByID_NotFound(t *testing.T) {
	mockService := new(MockUserService)
	mockService.On("GetUserByID", 1).Return(nil, fmt.Errorf("user not found"))
	userHandler := handler.NewUserHandler(mockService)

	router := mux.NewRouter()
	router.HandleFunc("/users/{id}", userHandler.GetUserByID).Methods("GET")

	req, err := http.NewRequest("GET", "/users/1", nil)
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	router.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusNotFound, rr.Code)
	var errorResponse helpers.ErrorResponse
	err = json.Unmarshal(rr.Body.Bytes(), &errorResponse)
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, "User not found", errorResponse.Error)
	assert.Equal(t, "The user with the specified ID does not exist", errorResponse.Details)
	mockService.AssertExpectations(t)
}

// TestUserHandler_CreateUser tests the CreateUser handler
func TestUserHandler_CreateUser(t *testing.T) {
	mockService := new(MockUserService)
	mockService.On("CreateUser", mock.Anything).Return(nil)
	userHandler := handler.NewUserHandler(mockService)

	user := model.User{Name: "New User", Email: "new@example.com"}
	userJSON, err := json.Marshal(user)
	if err != nil {
		t.Fatal(err)
	}
	req, err := http.NewRequest("POST", "/users", bytes.NewBuffer(userJSON))
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")
	rr := httptest.NewRecorder()
	userHandler.CreateUser(rr, req)

	assert.Equal(t, http.StatusCreated, rr.Code)
	var response map[string]string
	err = json.Unmarshal(rr.Body.Bytes(), &response)
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, "User created successfully", response["message"])
}

// TestUserHandler_UpdateUser tests the UpdateUser handler
func TestUserHandler_UpdateUser(t *testing.T) {
	mockService := new(MockUserService)
	mockService.On("GetUserByID", 1).Return(&model.User{ID: 1, Name: "Old User", Email: "old@example.com"}, nil)
	mockService.On("UpdateUser", mock.Anything).Return(nil)
	userHandler := handler.NewUserHandler(mockService)

	router := mux.NewRouter()
	router.HandleFunc("/users/{id}", userHandler.UpdateUser).Methods("PUT")

	updatedUser := model.User{ID: 1, Name: "Updated User", Email: "updated@example.com"}
	userJSON, err := json.Marshal(updatedUser)
	if err != nil {
		t.Fatal(err)
	}
	req, err := http.NewRequest("PUT", "/users/1", bytes.NewBuffer(userJSON))
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")
	rr := httptest.NewRecorder()
	router.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)
	var response map[string]string
	err = json.Unmarshal(rr.Body.Bytes(), &response)
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, "User updated successfully", response["message"])
	mockService.AssertExpectations(t)
}

// TestUserHandler_DeleteUser tests the DeleteUser handler with valid ID
func TestUserHandler_DeleteUser_ValidID(t *testing.T) {
	mockService := new(MockUserService)
	mockService.On("GetUserByID", 1).Return(&model.User{ID: 1, Name: "Test User", Email: "test@example.com"}, nil)
	mockService.On("DeleteUser", 1).Return(nil)
	userHandler := handler.NewUserHandler(mockService)

	router := mux.NewRouter()
	router.HandleFunc("/users/{id}", userHandler.DeleteUser).Methods("DELETE")

	req, err := http.NewRequest("DELETE", "/users/1", nil)
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	router.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)
	var response map[string]string
	err = json.Unmarshal(rr.Body.Bytes(), &response)
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, "User deleted successfully", response["message"])
	mockService.AssertExpectations(t)
}

// TestUserHandler_DeleteUser_NotFound tests the DeleteUser handler with invalid ID
func TestUserHandler_DeleteUser_NotFound(t *testing.T) {
	mockService := new(MockUserService)
	mockService.On("GetUserByID", 1).Return(nil, fmt.Errorf("user not found"))
	userHandler := handler.NewUserHandler(mockService)

	router := mux.NewRouter()
	router.HandleFunc("/users/{id}", userHandler.DeleteUser).Methods("DELETE")

	req, err := http.NewRequest("DELETE", "/users/1", nil)
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	router.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusNotFound, rr.Code)
	var errorResponse helpers.ErrorResponse
	err = json.Unmarshal(rr.Body.Bytes(), &errorResponse)
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, "User not found", errorResponse.Error)
	assert.Equal(t, "The user with the specified ID does not exist", errorResponse.Details)
	mockService.AssertExpectations(t)
}
