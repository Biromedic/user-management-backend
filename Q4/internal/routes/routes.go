package routes

import (
	"Q4/internal/handler"
	"Q4/internal/repository"
	"Q4/internal/service"
	"database/sql"
	"github.com/gorilla/mux"
)

func SetupRouter(db *sql.DB) *mux.Router {
	repo := repository.NewSQLUserRepository(db)
	services := service.NewUserService(repo)
	handlers := handler.NewUserHandler(services)

	router := mux.NewRouter()

	router.HandleFunc("/users", handlers.GetAllUsers).Methods("GET")
	router.HandleFunc("/users/{id}", handlers.GetUserByID).Methods("GET")
	router.HandleFunc("/users", handlers.CreateUser).Methods("POST")
	router.HandleFunc("/users/{id}", handlers.UpdateUser).Methods("PUT")
	router.HandleFunc("/users/{id}", handlers.DeleteUser).Methods("DELETE")

	return router
}
