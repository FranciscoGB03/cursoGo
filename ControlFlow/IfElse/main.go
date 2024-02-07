package main

import (
	"fmt"
	"math/rand"
)

func main() {
	//Basic if statement: The basic if statement is used to
	//execute a block of code based on the value of a Boolean expression

	age := 18
	if age >= 18 {
		fmt.Println("You are an adult")
	} else {
		fmt.Println("You are a minor")
	}

	//If statement with initialization: the if statement can
	//also, have an initialization statement executed
	//before evaluating the Boolean expression
	if num := rand.Intn(10); num > 5 {
		fmt.Println(num, "is greater than 5")
	} else {
		fmt.Println(num, "is less than or equal than 5")
	}

	//If statement with short-circuit evaluation: the if
	// a statement can use the short-circuit evaluation to
	//evaluate the second Boolean expression only
	//if the first Boolean expression is true

	x, y := 10, 50
	if x > 0 && y/x > 2 {
		fmt.Println(x, y)
	}

	//If statement with multiple conditions: the if
	// The statement can evaluate multiple Boolean
	// expressions separated by logical operators
	// (&& or ||)

	if age >= 18 && age <= 30 {
		fmt.Println("You are a young adult")
	} else if age > 30 && age <= 50 {
		fmt.Println("You are in your prime")
	} else {
		fmt.Println("You are either too young or too old")
	}

	// If statement with optional else if and else blocks:
	// The if statement can have multiple else if blocks
	// and an optional else block

	if x > 0 {
		fmt.Println("x is positive")
	} else if x == 0 {
		fmt.Println("x is zero")
	} else {
		fmt.Println("x is negative")
	}
}
