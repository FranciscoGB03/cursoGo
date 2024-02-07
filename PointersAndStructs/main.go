package main

import "fmt"

type Person struct {
	name string
	age  int
}

func changeName(p *Person, newName string) {
	p.name = newName
}

func main() {
	x := 5
	y := &x
	fmt.Println("x:", x)
	fmt.Println("y:", *y)

	*y = 10
	fmt.Println("x: ", x)
	fmt.Println("y: ", *y)

	p := Person{
		name: "Alice",
		age:  25,
	}
	fmt.Println(p)
	fmt.Println(p.name)
	fmt.Println(p.age)

	changeName(&p, "Bob")
	fmt.Println(p)
}
