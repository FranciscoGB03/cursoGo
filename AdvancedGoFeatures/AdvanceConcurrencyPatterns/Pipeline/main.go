package main

import "fmt"

func generator(nums ...int) <-chan int {
	out := make(chan int)
	go func() {
		for _, n := range nums {
			out <- n
		}
		close(out)
	}()
	return out
}

func square(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		for n := range in {
			out <- n * n
		}
		close(out)
	}()
	return out
}

func sum(in <-chan int) int {
	sum := 0
	for n := range in {
		sum += n
	}
	return sum
}

func main() {
	//create a pipeline of stages: generator->square->sum
	nums := []int{1, 2, 3, 4, 5}
	gen := generator(nums...)
	sq := square(gen)
	result := sum(sq)
	fmt.Println("Result:", result)
}
