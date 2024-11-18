package db

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/lib/pq"
)

func ConnectDB() (*sql.DB, error) {
	connStr := os.Getenv("DATABASE_URL")
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatalf("failed to open database: %v", err)
		return nil, err
	}
	if err := db.Ping(); err != nil {
		log.Fatalf("failed to ping database: %v", err)
		return nil, err
	}
	return db, nil
}
