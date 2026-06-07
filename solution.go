package main

import (
	"fmt"
	"sync"
)

func main() {
	// var mu sync.Mutex
	var wg sync.WaitGroup

	p := 0
	for i := 0; i < 5; i++ {
		wg.Add(1)
		// go increment(&p, &mu, &wg)
		go increment(&p, &wg)
	}

	wg.Wait()
	fmt.Println("final value: ", p)
}

// func increment(i *int, mu *sync.Mutex, wg *sync.WaitGroup) {
func increment(i *int, wg *sync.WaitGroup) {
	// mu.Lock()
	defer wg.Done()
	// defer mu.Unlock()
	*i++
}
