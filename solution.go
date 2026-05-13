package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func main() {
	var wg sync.WaitGroup
	var counter int64

	for i := 0; i < 100; i++ {
		wg.Add(1)
		go task(&wg, &counter)
	}

	wg.Wait()
	fmt.Println("final counter: ", atomic.LoadInt64(&counter))
}

func task(wg *sync.WaitGroup, counter *int64) {
	defer wg.Done()
	atomic.AddInt64(counter, 1)
}
