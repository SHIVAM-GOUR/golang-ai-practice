package main

import (
	"fmt"
)

func main() {
	ch1 := make(chan int, 1)
	ch2 := make(chan int, 1)
	ch3 := make(chan int, 1)

	go func() { ch1 <- 1 }()
	go func() { ch2 <- 2 }()
	go func() { ch3 <- 3 }()
	// ch1 <- 1
	// ch2 <- 2
	// ch3 <- 3

	select {
	case res := <-ch1:
		fmt.Println("res: ", res)
	case res := <-ch2:
		fmt.Println("res: ", res)
	case res := <-ch3:
		fmt.Println("res: ", res)
	default:
		fmt.Println("no channels ready")
	}

}
