package main

// Write your solution here.
// Challenge: G01 — Launch N goroutines, collect results with WaitGroup
//
// Problem:
// Write a function that launches N goroutines, each computing the square
// of its index (0..N-1). Collect all results into a slice in the correct
// order and return it. Use sync.WaitGroup for coordination.
//
// Requirements:
// - No goroutine should leak
// - Results must be in order [0², 1², 2², ..., (N-1)²]
// - Must compile clean with: go build -race ./...

import "fmt"

func squares(n int) []int {
	// your solution here
	return nil
}

func main() {
	result := squares(5)
	fmt.Println(result) // expected: [0 1 4 9 16]
}
