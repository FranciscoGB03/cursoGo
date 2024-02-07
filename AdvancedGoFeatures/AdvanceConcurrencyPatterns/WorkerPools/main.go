package main

import "fmt"

func worker(id int, jobs <-chan int, results chan<- int) {
	for j := range jobs {
		fmt.Println("worker", id, "processing job", j)
		results <- j * 2
	}
}

func main() {
	jobs := make(chan int, 100)
	results := make(chan int, 100)

	//create 3 workers
	for w := 1; w <= 3; w++ {
		go worker(w, jobs, results)
	}

	//add 9 jobs to the queue
	for i := 1; i <= 9; i++ {
		jobs <- i
	}

	close(jobs)

	// collect results from the workers
	for a := 1; a <= 9; a++ {
		<-results
	}
}
