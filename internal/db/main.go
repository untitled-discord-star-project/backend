package db

import (
	"errors"
)

type Cat struct {
	Name string
	Age  int
}

func Select() ([]Cat, error) {
	// For demonstration, let's say an error can occur here
	if false { // replace this with actual error condition
		return nil, errors.New("an error occurred")
	}

	return []Cat{
		{Name: "Fluffy", Age: 3},
		{Name: "Whiskers", Age: 5},
		{Name: "Mittens", Age: 2},
		{Name: "Snowball", Age: 7},
		{Name: "Mr. Bigglesworth", Age: 1},
		{Name: "Garfield", Age: 4},
		{Name: "Tom", Age: 6},
		{Name: "Sylvester", Age: 8},
	}, nil
}
