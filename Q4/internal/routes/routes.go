package routes

import (
	"Q4/config"
	"Q4/internal/handler"
	"Q4/internal/repository"
	"Q4/internal/service"
	"database/sql"
	"github.com/gorilla/mux"
	httpSwagger "github.com/swaggo/http-swagger"
)

func SetupRouter(db *sql.DB) *mux.Router {
	repo := repository.NewSQLUserRepository(db)
	services := service.NewUserService(repo)
	handlers := handler.NewUserHandler(services)

	router := mux.NewRouter()

	apiRouter := router.PathPrefix("/api/v1").Subrouter()
	apiRouter.Use(config.CorsMiddleware)

	apiRouter.HandleFunc("/users", handlers.GetAllUsers).Methods("GET")
	apiRouter.HandleFunc("/users/{id}", handlers.GetUserByID).Methods("GET")
	apiRouter.HandleFunc("/users", handlers.CreateUser).Methods("POST")
	apiRouter.HandleFunc("/users/{id}", handlers.UpdateUser).Methods("PUT")
	apiRouter.HandleFunc("/users/{id}", handlers.DeleteUser).Methods("DELETE")

	router.PathPrefix("/swagger/").Handler(httpSwagger.Handler(
		httpSwagger.URL("/swagger/doc.json")))

	return router
}
