package main

import "fmt"

// This function takes two parameters, both of which are integers,
// and returns an integer

func add(x int, y int) int {
	return x + y
}

// Functions can also return multiple values, like this:
func addAndMultiply(x int, y int) (int, int) {
	return x + y, x * y
}

// Functions can also have named return values, like this:
func addAndMultiplyNew(x int, y int) (addResult int, mulResult int) {
	addResult = x + y
	mulResult = x * y
	return
}

// It's also possible to use variadic functions, a function that can
// take a variable number of arguments. You can use the ... to indicate
// that a function can take a variable number of arguments of a certain
// type.
func sum(numbers ...int) int {
	sum := 0
	for _, num := range numbers {
		sum += num
	}
	return sum
}

func main() {
	result := add(1, 2)

	fmt.Println(result)

	addResult, mulResult := addAndMultiply(1, 2)
	fmt.Println(addResult, mulResult)

	_, multiplyResult := addAndMultiply(1, 2)
	fmt.Println(multiplyResult)

	addResult, mulResult = addAndMultiplyNew(1, 2)
	fmt.Println(addResult, mulResult)

	result = sum(1, 2, 3, 4, 5, 6, 7, 8)
	fmt.Println(result)
}
