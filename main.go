package main

import (
	"log"
	"os"
	"time"
	"github.com/joho/godotenv"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"database/sql"
)


func migration() {
	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file found or failed to load")
	}

	// this is the different URL for local development or dev from the container

	dbURL := os.Getenv("AmoradbConnectionString")
	//dbURL := os.Getenv("AmoradbConnectionStringLocal")
	if dbURL == "" {
		log.Fatal("AmoradbConnectionString is not set in environment")
	}

	maxRetries := 10
	for i := 0; i < maxRetries; i++ {
		db, err := sql.Open("postgres", dbURL)
		if err == nil {
			err = db.Ping()
			if err == nil {
				log.Println("Database is ready")
				break
			}
		}
		log.Printf("Waiting for database... (%d/%d)\n", i+1, maxRetries)
		time.Sleep(3 * time.Second)
	}

	m, err := migrate.New("file://db/migrations", dbURL)
	if err != nil {
		log.Fatalf("Failed to create migrate instance: %v", err)
	}
	if err := m.Up(); err != nil && err.Error() != "no change" {
		log.Fatalf("Migration failed: %v", err)
	}

	log.Println("Migration completed successfully.")
}


func main() {
	migration()
}

