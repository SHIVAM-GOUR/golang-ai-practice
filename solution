package main

import (
	"log"
	"sync"
)

func main() {
	var wg sync.WaitGroup

	wg.Add(1)
	go print("hello world", &wg)

	wg.Wait()
}

func print(name string, wg *sync.WaitGroup) {
	defer wg.Done()
	log.Println(name)
}
