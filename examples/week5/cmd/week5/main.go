package main

import (
	"week5/internal/infrastructure"
	"week5/internal/server"
)

func main() {
	api := infrastructure.New()
	s := server.New(":8080", api)
	s.Start()
}
