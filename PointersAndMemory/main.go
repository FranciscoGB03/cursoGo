package main

import "fmt"

func main() {
	p := new(int)   //allocate memory for an int and return a pointer
	fmt.Println(*p) // print the value of the int (which is 0, the zero value of int)
	*p = 2          //set the value of the int through the pointer
	fmt.Println(*p) //print the value of the int (which is now 2)

	s := make([]int, 0, 10)     //allocate memory for a slice of integers with a length of 0 and capacity of 10
	fmt.Println(len(s), cap(s)) // print the length and capacity of the slice (0 and 10, respectively)
	s = append(s, 1)            // add an element to the slice
	fmt.Println(len(s), cap(s)) // print the length and capacity of the slice(1 and 10, respectively)

}
