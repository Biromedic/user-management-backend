package handler

import (
	"Q4/internal/helpers"
	"Q4/internal/model"
	"Q4/internal/service"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"io"
	"net/http"
	"strconv"
)

type UserHandler struct {
	Service service.UserServiceInterface
}

func NewUserHandler(service service.UserServiceInterface) *UserHandler {
	return &UserHandler{
		Service: service,
	}
}

type ErrorResponse struct {
	Message string `json:"message"`
}

func writeErrorResponse(w http.ResponseWriter, status int, message string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	err := json.NewEncoder(w).Encode(ErrorResponse{Message: message})
	if err != nil {
		return
	}
}

func (uh *UserHandler) GetAllUsers(rw http.ResponseWriter, r *http.Request) {
	users, err := uh.Service.GetAllUsers()
	if err != nil {
		logrus.Errorf("Failed to retrieve users: %v", err)
		http.Error(rw, "Failed to retrieve users", http.StatusInternalServerError)
		return
	}

	rw.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(rw).Encode(users)
	if err != nil {
		logrus.Errorf("Failed to encode users: %v", err)
		return
	}
	logrus.Info("Users retrieved successfully")
}

func (uh *UserHandler) GetUserByID(rw http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	idStr, ok := params["id"]
	if !ok || idStr == "" {
		logrus.Warn("User ID is missing in request")
		helpers.WriteErrorResponse(rw, http.StatusBadRequest, "User ID is missing", "No 'id' parameter found in the URL")
		return
	}

	id, err := strconv.Atoi(idStr)
	if err != nil {
		logrus.Warnf("Invalid user ID: %s", idStr)
		helpers.WriteErrorResponse(rw, http.StatusBadRequest, "Invalid user ID", "The provided ID must be a numeric value")
		return
	}

	user, err := uh.Service.GetUserByID(id)
	if err != nil {
		logrus.Errorf("Failed to retrieve user with ID %d: %v", id, err)
		helpers.WriteErrorResponse(rw, http.StatusNotFound,
			"User not found",
			"The user with the specified ID does not exist")
		return
	}

	logrus.Infof("User with ID %d retrieved successfully", id)
	rw.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(rw).Encode(user); err != nil {
		logrus.Errorf("Failed to encode user with ID %d: %v", id, err)
		helpers.WriteErrorResponse(rw, http.StatusInternalServerError, "Failed to encode user", err.Error())
	}
}

func (uh *UserHandler) CreateUser(rw http.ResponseWriter, r *http.Request) {
	var user model.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		logrus.Warn("Invalid user data provided")
		helpers.WriteErrorResponse(rw, http.StatusBadRequest, "Invalid user data", "The request body must be valid JSON")
		return
	}

	err = uh.Service.CreateUser(&user)
	if err != nil {
		logrus.Errorf("Failed to create user: %v", err)
		helpers.WriteErrorResponse(rw, http.StatusInternalServerError, "Failed to create user", err.Error())
		return
	}

	rw.Header().Set("Content-Type", "application/json")
	rw.WriteHeader(http.StatusCreated)
	err = json.NewEncoder(rw).Encode(map[string]string{
		"message": "User created successfully",
	})
	if err != nil {
		logrus.Errorf("Failed to encode create user response: %v", err)
		return
	}
	logrus.Info("User created successfully")
}

func (uh *UserHandler) UpdateUser(rw http.ResponseWriter, r *http.Request) {
	id, err := getUserIDFromURL(r)
	if err != nil {
		helpers.WriteErrorResponse(rw, http.StatusBadRequest, "Invalid user ID", err.Error())
		logrus.Warn(err.Error())
		return
	}

	user, err := decodeUserFromBody(r.Body)
	if err != nil {
		helpers.WriteErrorResponse(rw, http.StatusBadRequest, "Invalid user data", err.Error())
		logrus.Warn(err.Error())
		return
	}

	user.ID = id

	if _, err := uh.Service.GetUserByID(user.ID); err != nil {
		logrus.Warnf("User with ID %d not found for update", user.ID)
		helpers.WriteErrorResponse(rw, http.StatusNotFound, "User not found", "The user with the specified ID does not exist")
		return
	}

	if err := uh.Service.UpdateUser(&user); err != nil {
		logrus.Errorf("Failed to update user with ID %d: %v", user.ID, err)
		helpers.WriteErrorResponse(rw, http.StatusInternalServerError, "Failed to update user", err.Error())
		return
	}

	respondWithSuccess(rw, "User updated successfully")
	logrus.Infof("User with ID %d updated successfully", user.ID)
}

func (uh *UserHandler) DeleteUser(rw http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	idStr, ok := params["id"]
	if !ok || idStr == "" {
		logrus.Warn("User ID is missing in delete request")
		helpers.WriteErrorResponse(rw, http.StatusBadRequest,
			"User ID is missing",
			"No 'id' parameter found in the URL")
		return
	}

	id, err := strconv.Atoi(idStr)
	if err != nil {
		logrus.Warnf("Invalid user ID: %s", idStr)
		helpers.WriteErrorResponse(rw, http.StatusBadRequest,
			"Invalid user ID",
			"The provided ID must be a numeric value")
		return
	}

	_, err = uh.Service.GetUserByID(id)
	if err != nil {
		logrus.Warnf("User with ID %d not found for deletion", id)
		helpers.WriteErrorResponse(rw, http.StatusNotFound,
			"User not found",
			"The user with the specified ID does not exist")
		return
	}

	err = uh.Service.DeleteUser(id)
	if err != nil {
		logrus.Errorf("Failed to delete user with ID %d: %v", id, err)
		helpers.WriteErrorResponse(rw, http.StatusInternalServerError, "Failed to delete user", err.Error())
		return
	}

	rw.Header().Set("Content-Type", "application/json")
	rw.WriteHeader(http.StatusOK)
	err = json.NewEncoder(rw).Encode(map[string]string{
		"message": "User deleted successfully",
	})
	if err != nil {
		logrus.Errorf("Failed to encode delete user response: %v", err)
		return
	}
	logrus.Infof("User with ID %d deleted successfully", id)
}

// Helper functions
func getUserIDFromURL(r *http.Request) (int, error) {
	vars := mux.Vars(r)
	idStr, ok := vars["id"]
	if !ok {
		return 0, fmt.Errorf("user ID not provided in URL")
	}

	id, err := strconv.Atoi(idStr)
	if err != nil || id <= 0 {
		return 0, fmt.Errorf("invalid user ID: %s", idStr)
	}

	return id, nil
}

func decodeUserFromBody(body io.ReadCloser) (model.User, error) {
	defer func(body io.ReadCloser) {
		err := body.Close()
		if err != nil {
			logrus.Warnf("Failed to close request body: %v", err)
		}
	}(body)
	var user model.User
	if err := json.NewDecoder(body).Decode(&user); err != nil {
		return user, fmt.Errorf("the request body must be valid JSON")
	}
	return user, nil
}

func respondWithSuccess(rw http.ResponseWriter, message string) {
	rw.Header().Set("Content-Type", "application/json")
	rw.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(rw).Encode(map[string]string{"message": message})
}
