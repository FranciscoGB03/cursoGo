package main

import (
	"context"
	"fmt"
	"time"
)

func longRunningOperation(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("Operation cancelled!")
			return
		default:
			fmt.Println("Operation running...")
			time.Sleep(time.Second)
		}
	}
}

func main() {
	// create a context with a cancellation signal
	ctx, cancel := context.WithCancel(context.Background())

	// start the long running operation in a goroutine
	go longRunningOperation(ctx)

	// wait for some time
	time.Sleep(5 * time.Second)

	// cancel the long-running operation
	cancel()

	// wait for the operation to terminate
	time.Sleep(time.Second)
	fmt.Println("Main thread exiting...")
}
