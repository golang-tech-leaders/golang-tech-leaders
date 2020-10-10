package main

import (
	"log"

	"demo/internal/database"
	"demo/internal/server"
)

const address = ":1234"

func main() {
	storage := database.NewRuntimeStorage()
	srv := server.New(storage)
	log.Println("starting server")
	err := srv.Run(address)
	if err != nil {
		log.Fatal(err)
	}
}
