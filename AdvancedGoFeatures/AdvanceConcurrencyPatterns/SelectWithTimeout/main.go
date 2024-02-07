package main

import (
	"fmt"
	"time"
)

func slowOperation() int {
	time.Sleep(3 * time.Second)
	return 42
}

func main() {
	c := make(chan int)

	// start a slow operation in a separate goroutine
	go func() {
		result := slowOperation()
		c <- result
	}()

	// wait for the slow operation with a timeout
	select {
	case result := <-c:
		fmt.Println("Result:", result)
	case <-time.After(2 * time.Second):
		fmt.Println("Timeout occurred!")
	}
}
