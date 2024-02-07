package main

import (
	"fmt"
	"time"
)

func worker(ch chan int) {
	x := <-ch
	fmt.Println("Received:", x)
}

func main() {
	ch := make(chan int)
	go worker(ch)
	ch <- 42
	fmt.Println("Sent:", 42)
	time.Sleep(1 * time.Millisecond)
}
