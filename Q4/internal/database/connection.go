package database

import (
	"database/sql"
	"log"

	_ "modernc.org/sqlite"
)

func NewConnection() *sql.DB {
	db, err := sql.Open("sqlite", "./users.db")
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	createTableQuery := `
	CREATE TABLE IF NOT EXISTS users (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT NOT NULL,
		email TEXT NOT NULL UNIQUE
	);
	`

	if _, err = db.Exec(createTableQuery); err != nil {
		log.Fatalf("Failed to create table: %v", err)
	}

	log.Println("Database connection established and table verified.")
	return db
}
