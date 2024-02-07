package main

import "fmt"

func main() {
	//var slice []type
	//slice:=make([]type,length, capacity)

	//Creating a slice
	// var slice []int //create an empty slice of int
	slice := make([]int, 3, 5) //create a slice with length 3 and capacity 5
	fmt.Println(slice)

	// Accessing elements
	newSlice := []int{1, 2, 3}
	fmt.Println(newSlice[0])  // prints 1
	fmt.Println(newSlice[1:]) // prints [2,3]

	//Modifying elements
	newSlice[0] = 4 //modifies the first element to be 4

	//Appending elements
	newSlice = append(newSlice, 5) //appends 5 to the end of the slice
	fmt.Println(newSlice)

	//Copying slices
	copySlice := make([]int, len(newSlice))
	copy(copySlice, newSlice) // copies the contents of newSlice into copySlice

	fmt.Println(copySlice)

	// Using a for loop with indexing
	for i := 0; i < len(newSlice); i++ {
		fmt.Printf("Index: %d, Value: %d\n", i, newSlice[i])
	}

	// Using a for loop with range
	for index, value := range newSlice {
		fmt.Printf("Index: %d, Value: %d\n", index, value)
	}

	//Using a for loop without index
	for _, value := range newSlice {
		fmt.Println(value)
	}

	//Using a for loop with index
	i := 0
	for i < len(newSlice) {
		fmt.Printf("Index: %d, Value: %d\n", i, newSlice[i])
		i++
	}

}
