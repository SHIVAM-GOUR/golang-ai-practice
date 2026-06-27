package main

import (
	"fmt"
	"sync"
)

func main() {
	ch := make(chan int, 3)
	var once sync.Once

	closeChannel := func() {
		once.Do(func() {
			close(ch)
		})
	}

	ch <- 1
	ch <- 2
	ch <- 3
	go closeChannel()
	go closeChannel()
	// go func() {
	// closeChannel()
	// }()
	// go func() {
	// 	ch <- 3
	// 	// closeChannel()
	// }()

	for val := range ch {
		fmt.Println(val)
	}

}
