/*
Go starts  with the package and with its name
Go uses the keyword import to import packages
Go uses the keyword func to define functions
Go uses curly braces {} to define the beginning and the end of a block of code
Go uses the keyword var to define variables
Go uses the keyword return to return a value from a function
*/

package main

import (
	"fmt"
)

func displayMessage() string {
	var msg string = "Hello world...!"
	return msg
}

func main() {
	fmt.Println(displayMessage())
}
