package main

import "fmt"

// Define a simple interface for a vehicle
type Vehicle interface {
	Drive() string
}

// Define a car type that implements the Vehicle interface
type Car struct {
	make  string
	model string
}

func (c Car) Drive() string {
	return fmt.Sprintf("Driving a %s %s", c.make, c.model)
}

// Define a truck type that also implements the Vehicle interface
type Truck struct {
	make  string
	model string
}

func (t Truck) Drive() string {
	return fmt.Sprintf("Driving a %s %s truck", t.make, t.model)
}

func main() {
	// Create a slice of vehicles that can hold both cars and trucks
	vehicles := []Vehicle{
		Car{"Toyota", "Camry"},
		Truck{"Ford", "F150"},
	}

	// Loop through the vehicles and call the Drive method on each one
	for _, v := range vehicles {
		fmt.Println(v.Drive())
	}
}
