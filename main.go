package main

import (
  "Amora/db/migration"
	"Amora/internal/Logging"
	"Amora/Init"
)

func main() {
	migration.Migrations(true);
	db := Init.DBinit()
	Logging.InitLogger(db)
	Init.InitOpenAi()
	Init.RoutesInit(db)
}


