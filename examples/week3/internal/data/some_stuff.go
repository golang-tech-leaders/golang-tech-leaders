package data

import (
	"errors"
	"fmt"
	"math/rand"

	"week3/internal/repository"
)

var (
	// ErrNotFournd not found
	ErrNotFournd = errors.New("not found")
	// ErrConnectionInterrupted connection is interrupted
	ErrConnectionInterrupted = errors.New("connection is interrupted")
)

// GetSomeData returns some data by key
func GetSomeData(key int) (string, error) {
	if err := checkConnection(); err != nil {
		return "", ErrConnectionInterrupted
	}

	val, err := repository.GetData(key)
	if errors.Is(err, repository.ErrNotFound) {
		return "", fmt.Errorf("%w: %v", ErrNotFournd, err)
	} else if err != nil {
		return "", fmt.Errorf("unknown error: %v", err)
	}
	return val, nil
}

// checkConnection emulates some random connection errors
func checkConnection() error {
	if rand.Intn(10) < 3 {
		return ErrConnectionInterrupted
	}
	return nil
}
