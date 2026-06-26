package main

import (
	"errors"
	"fmt"
)

var ErrNotFound = errors.New("not found")

func main() {

	err := checkErr()
	if errors.Is(err, ErrNotFound) {
		fmt.Println("Error Matched")
	} else {
		fmt.Println("Error not matched")
	}

}

func checkErr() error {
	return ErrNotFound
}
