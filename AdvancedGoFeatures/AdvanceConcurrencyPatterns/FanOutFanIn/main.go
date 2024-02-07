package main

import (
	"fmt"
	"sync"
)

func main() {
	input := []int{1, 2, 3, 4, 5}

	//fan-out
	chans := make([]chan int, 3)
	for i := 0; i < 3; i++ {
		chans[i] = make(chan int)
		go square(input, chans[i])
	}

	//fan-in
	var wg sync.WaitGroup
	wg.Add(len(chans))
	out := make(chan int)
	for _, c := range chans {
		go func(c <-chan int) {
			defer wg.Done()
			for v := range c {
				out <- v
			}
		}(c)
	}

	go func() {
		wg.Wait()
		close(out)
	}()

	// print results
	for res := range out {
		fmt.Println(res)
	}
}

func square(input []int, out chan<- int) {
	for _, i := range input {
		out <- i * i
	}
	close(out)
}
