package main

import "fmt"

var (
	x int
	y float32
	z string
)

func main() {
	//type of valid variables
	var speed int = 1
	var Speed int = 2
	var _speed int = 3

	fmt.Println(speed, Speed, _speed)
	//Declare a single variable and initialize
	var a = 5
	//You can also declarate multiple variables at once
	var b, c, d = 5, 10, 15
	//You can also leave the the type and initialize the variable
	// with its value(Using the shorthand notation with the :=)
	e, f, g := 20, 25, 30
	fmt.Println(a, b, c, d, e, f, g, x, y, z)
}
