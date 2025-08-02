package Init


import (
	"time"
	"os"
	"log"
	"gorm.io/gorm"
	"gorm.io/driver/postgres"
	"Amora/internal/Logging"
	"github.com/joho/godotenv"
)



func DBinit() *gorm.DB {
    // load env vars
    err := godotenv.Load()
    if err != nil {
        log.Printf("Warning: failed to load .env file: %v", err)
    }

    dbURL := os.Getenv("AmoradbConnectionStringLocal")
    var db *gorm.DB

    const maxRetries = 10

    for i := 0; i < maxRetries; i++ {
        db, err = gorm.Open(postgres.Open(dbURL), &gorm.Config{})
        if err == nil {
            sqlDB, err := db.DB()
            if err == nil && sqlDB.Ping() == nil {
                log.Println("Database connection established")
                return db
            }
        }

        log.Printf("Failed to connect to database, retrying... (%d/%d): %v", i+1, maxRetries, err)
        time.Sleep(3 * time.Second)
    }

    // If we reach here, all retries failed
    Logging.Logs.Fatalf("Failed to connect to database after %d attempts: %v", maxRetries, err)
    return nil // won't reach here due to Fatalf but just to satisfy compiler
}

