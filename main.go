package main

import (
	"log"
	"payments/cmd"
)

func main() {
	if err := cmd.RunApp(); err != nil {
		log.Fatal(err)
	}
	log.Println("done")
}
