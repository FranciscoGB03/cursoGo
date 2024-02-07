package main

import "fmt"

func main() {
	// Basic for loop
	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}

	// for loop with a range
	numbers := []int{2, 3, 4, 5, 6, 7, 8}
	for i, num := range numbers {
		fmt.Printf("index: %d, value: %d\n", i, num)
	}

	//goto loop sintax
	i := 0
loop:
	fmt.Println(i)
	i++
	if i < 5 {
		goto loop
	}

	//for loop with a single condition
	i = 0
	for i < 5 {
		fmt.Println(i)
		i++
	}

	// for loop with break and continue
	for i := 0; i < 10; i++ {
		if i == 5 {
			continue //skip the current iteration
		}
		if i == 8 {
			break //exit the loop
		}
		fmt.Println(i)
	}

	//Infinite for loop
	/*for {
		fmt.Println("I will run forever")
	}*/

}
