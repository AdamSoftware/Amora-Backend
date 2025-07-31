package main

import (
  "Amora/db/migration"
	"Amora/internal/Logging"
	"Amora/Init"
)

func main() {
	// run the migration unless it is turned off
	migration.Migrations(false);
	Logging.InitLogger()
	// init all the database connections and setup
	db := Init.DBinit()

	// initialize the routes
	Init.RoutesInit(db)

	// attach the database to the logger
	Logging.Logs.AddHook(Logging.NewGormHook(db))
}


