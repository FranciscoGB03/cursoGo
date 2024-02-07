package main

import (
	"fmt"
	"time"
)

func worker1(done chan<- bool) {
	fmt.Println("Worker 1 working...")
	time.Sleep(2 * time.Second)
	done <- true
}

func worker2(done chan<- bool) {
	fmt.Println("Worker 2 working...")
	time.Sleep(4 * time.Second)
	done <- true
}

func main() {
	// create two channels to signal the completation of each worker
	done1 := make(chan bool, 1)
	done2 := make(chan bool, 1)
	// start the workers in separate goroutines
	go worker1(done1)
	go worker2(done2)

	// wait for both workers to complete using a select statement
	for i := 0; i < 2; {
		select {
		case <-done1:
			i++
		case <-done2:
			i++
		}
	}

	fmt.Println("Both workers completed!")
}
