package main

import (
	"fmt"
)

func main() {
	// err := work()
	// if err != nil {
	// 	fmt.Println(err)
	// }

	defer func() {
		r := recover()
		if r != nil {
			fmt.Println("caught: ", r)
		}

		go func() {
			panic("goroutine panic")
		}()
	}()

}

// func work() (err error) {
// 	defer func() {
// 		r := recover()
// 		if r != nil {
// 			err = fmt.Errorf("recovered: %v", r)
// 		}
// 	}()
// 	panic("panic occurred")
// 	fmt.Println("Task Scuccess")
// 	return nil
// }
