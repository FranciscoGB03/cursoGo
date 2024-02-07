package main

import "fmt"

func main() {
	var numbers = [5]int{1, 2, 3, 4, 5}
	fmt.Println(numbers[0])

	//modifying values
	numbers[0] = 6
	fmt.Println(numbers) // prints [6, 2, 3, 4, 5]

	var nums = [3][3]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}
	fmt.Println(nums)

	// you can iterate over arrays in GO using loops
	for i := 0; i < len(numbers); i++ {
		fmt.Println("Numbers index: ", i, "Value: ", numbers[i])
	}

	//You can traverse this array using a nested for loop
	for i := 0; i < len(nums); i++ {
		for j := 0; j < len(nums[i]); j++ {
			fmt.Print(nums[i][j], " ")
		}
		fmt.Println()
	}

	//You can also use the for loop with the range keyword to
	//iterate over the elements of an array:
	fmt.Println("Using range")
	for _, v := range nums {
		for _, k := range v {
			fmt.Print(k, " ")
		}
		fmt.Println()
	}
}
