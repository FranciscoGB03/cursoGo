package main

import "fmt"

type Shape interface {
	Area() float64
}

type Square struct {
	Length float64
}

func (s Square) Area() float64 {
	return s.Length * s.Length
}

type Circle struct {
	Radius float64
}

func (c Circle) Area() float64 {
	return 3.14 * c.Radius * c.Radius
}

func main() {
	shapes := []Shape{Square{Length: 5}, Circle{Radius: 3}}
	for _, shape := range shapes {
		fmt.Println(shape.Area())
	}
}
