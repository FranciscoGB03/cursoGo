package main

import (
	"fmt"
	"time"
)

func main() {
	// create a rate limiter that allows 2 operations per second
	limiter := time.Tick(time.Second / 2)
	// execute 6 operations, rate-limited to 2 per second
	for i := 0; i < 6; i++ {
		fmt.Println("Operation", i+1, "starting...")
		<-limiter
		fmt.Println("Operation", i+1, "completed!")

	}
	// execute 6 operations, rate-limited to 2 per second
	for i := 0; i < 6; i++ {
		fmt.Println("Operation", i+1, "starting...")
		<-limiter
		fmt.Println("Operation", i+1, "completed!")

	}
}
