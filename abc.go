package main

// import (
// 	"fmt"
// 	"sync"
// )

// func main() {
// 	counter := 0
// 	var mu sync.Mutex
// 	var wg sync.WaitGroup

// 	for i := 0; i < 5; i++ {
// 		wg.Add(1)
// 		go increment(&counter, &mu, &wg)
// 	}

// 	wg.Wait()
// 	fmt.Println("counter: ", counter)

// }

// func increment(n *int, mu *sync.Mutex, wg *sync.WaitGroup) {
// 	defer wg.Done()
// 	mu.Lock()
// 	defer mu.Unlock()
// 	for i := 0; i < 1000; i++ {
// 		*n++
// 	}
// }
