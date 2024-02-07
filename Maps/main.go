package main

import "fmt"

func main() {
	// Declare a map with string keys and int values
	var myMap map[string]int

	// Initialize the map with make
	myMap = make(map[string]int)

	// Add values to the map
	myMap["key1"] = 1
	myMap["key2"] = 2

	// Access a value in tha map
	value1 := myMap["key1"]
	fmt.Println(value1)

	// check if a key exists in the map
	value2, ok := myMap["key2"]
	if ok {
		// key exists, value 2 contains the value associated with "key2"
		fmt.Println(value2)
	} else {
		// key does not exist
		fmt.Println("key2 does not exist")
	}
	// Delete a key-value pair from the map
	delete(myMap, "key2")

	value2, ok = myMap["key2"]
	if ok {
		//key exists, value 2 contains the value associated with "key2"
		fmt.Println(value2)
	} else {
		//key does not exist
		fmt.Println("key2 does not exit")
	}

	myMap = map[string]int{"one": 1, "two": 2, "three": 3}
	//Using a for range
	for key, value := range myMap {
		fmt.Println(key, value)
	}
	//Using a for loop with map keys
	for key := range myMap {
		fmt.Println(key)
	}
	// Using for loop with map values
	for _, value := range myMap {
		fmt.Println(value)
	}
}
