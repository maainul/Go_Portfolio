package main

import (
	"log"
	"portfolio/api/storage/postgres"
)

func main() {
	if err := postgres.Migrate(); err != nil {
		log.Fatal(err)
	}
}
