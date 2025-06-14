package main

import (
	"errors"
	"fmt"
)

// EX01 - Create a sentinel error to represent an invalid ID.
// In `main`, use `errors.Is` to check for the sentinel error,
// and pirnt a message when it is found.
var ErrInvalidId error = errors.New("Invalid Id")

func main() {
	const x int = 42
	fmt.Printf("hey there")
}
