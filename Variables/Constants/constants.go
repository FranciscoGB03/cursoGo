package main

import "fmt"

// Using the const keyword
const pi float64 = 3.14159

// Using shorthand notation
const e = 2.718

// Declaring multiple constants in a single statement
const (
	Mon = 1
	Tue = 2
	Wed = 3
	Thu = 4
	Fri = 5
	Sat = 6
	Sun = 7
)

// It's also possible to use iota to declare a set of related constants, these are good for enumerations
const (
	Unknown = iota
	Female
	Male
)

func main() {
	fmt.Println(pi, e)
	fmt.Println(Mon, Tue, Wed, Thu, Fri, Sat, Sun)
	fmt.Println(Female, Male)
}
