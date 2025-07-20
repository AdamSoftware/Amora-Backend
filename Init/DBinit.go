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
	var db *gorm.DB

	// load the environment vars
	err := godotenv.Load()

	dbURL := os.Getenv("AmoradbConnectionString")
	//dbURL := os.Getenv("AmoradbConnectionStringLocal")

	// retry connection if it fails
  MaxRetries := 10
	for i := 0; i < MaxRetries; i++ {
		// attempt to connect the exit if nothing goes wrong
		if err == nil {
			break
		}
		// Log error of debugging
		log.Printf("Failed to connect to database, retrying... (%d/%d)", i+1, MaxRetries)
		db, err = gorm.Open(postgres.Open(dbURL), &gorm.Config{})

		// setting the sqlDB to the db connection
		sqlDB, err := db.DB()
		if err == nil && sqlDB.Ping() == nil {
				log.Println("Database is ready")
				break
		}
		log.Printf("Waiting for database... (%d/%d)\n", i+1, MaxRetries)
		time.Sleep(3 * time.Second)
		
	}

	// If connection is total failure, Log to the database and exit the connection
	if err != nil {
		Logging.Logs.Fatalf("Failed to connect to database: %v", err)
	}

	// return the database connection
	return db
}
