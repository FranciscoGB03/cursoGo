package main

import (
	"fmt"
	"time"
)

func longRunningOperation(cancel <-chan struct{}) {
	for {
		select {
		default:
			fmt.Println("Long-running operation running...")
			time.Sleep(time.Second)
		case <-cancel:
			fmt.Println("Long-running operation cancelled!")
			return
		}
	}
}

func main() {
	// create a cancellation channel
	cancel := make(chan struct{})

	// start the long-running operation in a goroutine
	go longRunningOperation(cancel)

	// wait for some time
	time.Sleep(5 * time.Second)

	// cancel the long-running operation
	close(cancel)

	// wait for the operation to terminate
	time.Sleep(time.Second)
	fmt.Println("Main thread exiting...")
}
