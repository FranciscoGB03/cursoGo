package main

import "fmt"

func main() {
	// we declare two anonymous functions add and subtract that take two
	// integer arguments and return an integer. We then pass these functions
	// as arguments to another functions apply that also takes two integers
	// and a function as arguments and returns an integer

	add := func(x, y int) int {
		return x + y
	}

	subtract := func(x, y int) int {
		return x - y
	}
	fmt.Println(apply(add, 2, 3))      // Output: 5
	fmt.Println(apply(subtract, 5, 2)) // Output:3
}

// The apply function applies the function f to the integer
// arguments x and y returns the result
func apply(f func(int, int) int, x, y int) int {
	return f(x, y)
}
