package main

import (
	"fmt"
)

func main() {
	ch := make(chan int)

	go func() {
		ch <- 1
		// ch <- 2
		// ch <- 3
	}()

	for i := 0; i < 3; i++ {
		fmt.Println(<-ch)
	}

}
