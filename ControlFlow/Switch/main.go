package main

import (
	"fmt"
	"math/rand"
)

//Type switch: The switch statement can be used to
//evaluate the type of interface value.

func PrintType(x interface{}) {
	switch x.(type) {
	case int:
		fmt.Println("x is an integer")
	case string:
		fmt.Println("x is a string")
	default:
		fmt.Println("x is of an unknown type")
	}
}

func main() {
	// Basic switch: The basic switch statement is use to
	// evaluate a variable or an expression and execute a
	// block of code based on the value of the variable.

	day := "Monday"
	switch day {
	case "Monday":
		fmt.Println("Today is Monday")
	case "Tuesday":
		fmt.Println("Today is Tuesday")
	case "Wednesday":
		fmt.Println("Today is Wednesday")
	case "Thursday":
		fmt.Println("Today is Thursday")
	case "Friday":
		fmt.Println("Today is Friday")
	case "Saturday":
		fmt.Println("Today is Saturday")
	case "Sunday":
		fmt.Println("Today is Sunday")
	default:
		fmt.Println("Value not valid")
	}

	//Switch with multiple expressions: The switch statement
	// can also evaluate multiple expressions separated by commas.

	x := 10

	switch x {
	case 5, 10:
		fmt.Println("x is 5 or 10")
	case 15, 20:
		fmt.Println("x is 15 or 20")
	default:
		fmt.Println("x is other than 5, 10, 15 or 20")
	}

	// Switch with initializer: The switch statement can also
	// hava an initializer statement executed before evaluating
	// the expressions.

	switch n := rand.Intn(10); {
	case n == 0:
		fmt.Println("It's zero")
	case n > 0 && n < 5:
		fmt.Println("It's a small number")
	case n >= 5 && n < 10:
		fmt.Println("It's a big number")
	}

	PrintType(x)

	// Switch with fallthrough: the fallthrough statement can
	// be used in a switch statement to execute the statements
	// of the next case block as well, regardless of the case
	// expression.

	x = 0
	switch x {
	case 0:
		fmt.Println("x is zero")
		fallthrough
	case 1:
		fmt.Println("x is less than or equal to 1")
		fallthrough
	case 2:
		fmt.Println("x is less than or equal to 2")
	default:
		fmt.Println("x is greater than 2")
	}
}
