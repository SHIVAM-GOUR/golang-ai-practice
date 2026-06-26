package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup

	wg.Add(1)
	go train(&wg)

	wg.Wait()
}

func train(wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("Hello from Goroutine")
}
