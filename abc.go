package main

import (
	"log"
)

func main() {
	// defer func() {
	// 	if r := recover(); r != nil {
	// 		log.Println("panic caught — send on closed channel:", r)
	// 	}
	// }()

	ch := make(chan int, 3)

	print(ch)

	close(ch);
	ch <- 5
	for v := range ch {
		log.Println(v)
	}

}

func print(ch chan int) {
	ch <- 1
	ch <- 2
	ch <- 3
}
