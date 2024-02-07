package main

import (
	"fmt"
	"time"
)

func hello() {
	fmt.Println("Hello, goroutine!")
}

func main() {
	go hello()
	fmt.Println("Main function")
	time.Sleep(1 * time.Millisecond)
}
