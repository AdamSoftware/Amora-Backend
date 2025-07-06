package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func main() {
	// Load the .env file
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	// Get the database connection string from the environment
	dbURL := os.Getenv("AmoradbConnectionString")
	if dbURL == "" {
		log.Fatal("DATABASE_URL is not set in .env")
	}

	// Create a new migration instance
	m, err := migrate.New(
		"file://db/migrations",
		dbURL,
	)
	if err != nil {
		log.Fatalf("Failed to create migrate instance: %v", err)
	}

	// Run the migrations
	if err := m.Up(); err != nil && err.Error() != "no change" {
		log.Fatalf("Migration failed: %v", err)
	}

	log.Println("Migration completed successfully.")
}

