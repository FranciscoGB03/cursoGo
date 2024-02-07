package main

import (
	"fmt"
	"sync"
)

func Swap(a, b int) (int, int) {
	return b, a
}

func main() {
	wg := sync.WaitGroup{}

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			fmt.Println(Swap(i+1, i))
		}(i)
	}
	wg.Wait()
}
