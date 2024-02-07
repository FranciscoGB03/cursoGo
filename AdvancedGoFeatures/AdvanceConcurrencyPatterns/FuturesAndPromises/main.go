package main

import "fmt"

func computeFuture() <-chan int {
	result := make(chan int, 1)
	go func() {
		// simulate a long running computation
		sum := 0

		for i := 0; i < 100000000; i++ {
			sum += i
		}
		result <- sum
	}()
	return result
}

func main() {
	// create a future and a promise
	future := computeFuture()
	promise := make(chan int, 1)

	//start a goroutine to wait for the future to complete
	go func() {
		promise <- <-future
	}()

	// do some other work in the main thread
	fmt.Println("Main thread doing some other work....")

	//wait for the future to complete and print the result
	fmt.Println("Future result:", <-promise)
}
