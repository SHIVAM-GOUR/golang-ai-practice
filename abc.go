package main

import (
	"fmt"
)

func abc() {
	i := 0
	defer fmt.Println(i)
	i = 1
}
