package main

import (
	"fmt"
	"reflect"
)

type Person struct {
	Name string
	Age  int
}

func main() {
	person := Person{"Fgb", 33}
	valueOfP := reflect.ValueOf(person)
	typeOfP := reflect.TypeOf(person)

	for i := 0; i < valueOfP.NumField(); i++ {
		fmt.Printf("Type of field %v and its value %v\n",
			typeOfP.Field(i), valueOfP.Field(i))
	}
}
