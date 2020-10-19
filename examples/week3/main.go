/*
Package main

Пример обработки ошибок в Go
*/
package main

import (
	"errors"
	"log"
	"math/rand"
	"time"

	"week3/internal/data"
)

func init() {
	// initialise randomizer with unique seed
	rand.Seed(time.Now().Unix())
}

func main() {
	val, err := data.GetSomeData(rand.Intn(10) - 3)
	// проверка, есть ли в цепочке ошибок ошибка data.ErrNotFournd
	if errors.Is(err, data.ErrNotFournd) {
		log.Fatalf("not found: %v", err)
		// проверка, является ли конкретная ошибка ошибкой data.ErrConnectionInterrupted
	} else if err == data.ErrConnectionInterrupted {
		log.Fatalf("connection error: %v", err)
		// общая проверка ошибки
	} else if err != nil {
		log.Fatalf("unknown error: %v", err)
	}
	log.Printf("value is %q", val)
}
