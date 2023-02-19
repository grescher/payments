package main

import (
	"log"
	"payments/db"
)

func main() {
	connDB, err := db.NewPostgresDB(db.DatabaseURL())
	if err != nil {
		log.Fatal(err)
	}
	defer connDB.Close()

	// r := repository.NewRepository(connDB)
	// s := service.NewService(r)
	// h := server.NewHandlers(s)
	// server.Run(h, config.ServerPort())
}
