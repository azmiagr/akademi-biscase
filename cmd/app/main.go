package main

import (
	"akademi-business-case/pkg/config"
	"akademi-business-case/pkg/database/mariadb"
	"log"
)

func main() {
	config.LoadEnvironment()

	db, err := mariadb.ConnectDatabase()
	if err != nil {
		log.Fatal(err)
	}

	err = mariadb.Migrate(db)
	if err != nil {
		log.Fatal(err)
	}
}
