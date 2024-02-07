package main

import "fmt"

type Animal interface {
	Name() string
	Speak() string
}

type Dog struct{}

func (d Dog) Name() string {
	return "Dog"
}

func (d Dog) Speak() string {
	return "Woof!"
}

type Cat struct{}

func (c Cat) Name() string {
	return "Cat"
}

func (c Cat) Speak() string {
	return "Meow!"
}

func main() {
	/**se genera un arreglo de animales que contiene dos: un perro y un gato*/
	animals := []Animal{Dog{}, Cat{}}
	/**se recorre el arreglo de animales*/
	for _, animal := range animals {
		fmt.Println(animal.Name(), "-->", animal.Speak())
	}
}
