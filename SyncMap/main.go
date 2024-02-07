package main

import (
	"fmt"
	"sync"
)

func main() {
	//create a new sync.Map
	var m sync.Map

	// Insert key-value pairs into the map
	m.Store("apple", 3)
	m.Store("banana", 5)
	m.Store("orange", 2)

	//Retrieve a value by key
	value, ok := m.Load("banana")
	if ok {
		fmt.Println("banana:", value)
	}

	// Update a value by key
	m.Store("banana", 6)

	// Delete a key-value pair
	m.Delete("orange")

	// Iterate over the keys and values in the map
	m.Range(func(key, value interface{}) bool {
		fmt.Println(key, value)
		return true
	})

}
