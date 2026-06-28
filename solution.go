package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func main() {
	var counter int64
	var wg sync.WaitGroup

	for i := 0; i < 100; i++ {
		wg.Add(1)
		go inc(&counter, 1, &wg)
	}

	wg.Wait()
	fmt.Println(counter)
}

func inc(counter *int64, new int64, wg *sync.WaitGroup) {
	defer wg.Done()
	atomic.AddInt64(counter, new)
}
