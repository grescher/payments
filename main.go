package main

import (
	"log"
	"payments/db"
)

func main() {
	storage, err := db.NewPostgresDB(db.DatabaseURL())
	check(err)
	defer storage.Close()

}

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
