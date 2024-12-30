package main

import (
	"Q4/internal/database"
	"Q4/internal/middleware"
	"Q4/internal/routes"
	"database/sql"
	"github.com/sirupsen/logrus"
	"log"
	"net/http"
)

func main() {
	logrus.SetFormatter(&logrus.TextFormatter{
		FullTimestamp: true,
	})
	logrus.SetLevel(logrus.InfoLevel)

	db := database.NewConnection()
	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(db)

	router := routes.SetupRouter(db)

	loggedRouter := middleware.LoggingMiddleware(router)

	log.Println("Server running on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", loggedRouter))
}
