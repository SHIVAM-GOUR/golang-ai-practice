package main

import (
	"fmt"
)

func main() {
	ch := make(chan int)
	ch2 := make(chan int)

	go demo(ch, 1)
	go demo(ch2, 2)

	select {
	case v := <-ch:
		fmt.Println(v)
	case v2 := <-ch2:
		fmt.Println(v2)
	}

}

func demo(ch chan int, value int) {
	ch <- value
}
