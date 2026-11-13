package Init


import (
	"time"
	"os"
	"log"
	"strconv"
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

		dbURL := os.Getenv("AmoradbConnectionString")
    //dbURL := os.Getenv("AmoradbConnectionStringLocal")
    // Wait for the database to accept connections before continuing.
    // In docker compose / Kubernetes the API can start before Postgres is ready,
    // so we keep retrying. Override the attempt count with DB_MAX_RETRIES.
    maxRetries := 20
    if v, convErr := strconv.Atoi(os.Getenv("DB_MAX_RETRIES")); convErr == nil && v > 0 {
        maxRetries = v
    }
    const retryDelay = 3 * time.Second

    for i := 1; i <= maxRetries; i++ {
        db, connErr := tryConnect(dbURL)
        if connErr == nil {
            log.Println("Database connection established")
            return db
        }
        err = connErr

        log.Printf("Waiting for database... (%d/%d): %v", i, maxRetries, err)
        time.Sleep(retryDelay)
    }

    // Use the DB-backed logger if it has been initialised; otherwise fall back
    // to the standard Go logger (on first startup Logging needs the DB to exist).
    if Logging.Logs != nil {
        Logging.Logs.Fatalf("Failed to connect to database after %d attempts: %v", maxRetries, err)
    }
    log.Fatalf("Failed to connect to database after %d attempts: %v", maxRetries, err)
    return nil // unreachable; log.Fatalf exits
}

// tryConnect opens a connection and verifies it with a ping.
// On failure it closes any half-open connection so retries don't leak.
func tryConnect(dbURL string) (*gorm.DB, error) {
    db, err := gorm.Open(postgres.Open(dbURL), &gorm.Config{})
    if err != nil {
        return nil, err
    }
    sqlDB, err := db.DB()
    if err != nil {
        return nil, err
    }
    if err := sqlDB.Ping(); err != nil {
        sqlDB.Close()
        return nil, err
    }
    return db, nil
}

