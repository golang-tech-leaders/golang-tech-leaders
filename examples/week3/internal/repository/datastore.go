package repository

import (
	"errors"
	"fmt"
)

var (
	// ErrNotFound data in not found
	ErrNotFound = errors.New("data not found in the database")
)

var database = map[int]string{
	0: "zero",
	1: "one",
	2: "two",
	3: "three",
}

// GetData return data from the database by a key
func GetData(key int) (string, error) {
	// some validation
	if key < 0 {
		return "", fmt.Errorf("invalid key: %d", key)
	}

	v, ok := database[key]
	if !ok {
		return "", ErrNotFound
	}
	return v, nil
}
