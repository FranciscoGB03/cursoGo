package main

import "fmt"

// Define a new struct implements the Printer interface
type FancyPrinter struct {
	Message string
}

func (fp FancyPrinter) Print() {
	fmt.Println("**********")
	fmt.Println(fp.Message)
	fmt.Println("**********")
}

// Define the printer interface that our objects depend on
type Printer interface {
	Print()
}

// Define a struct for our message printer that implements the Printer interface
type MessagePrinter struct {
	Message string
}

func (mp MessagePrinter) Print() {
	fmt.Println(mp.Message)
}

// Define a struct for our greeter that depends on a Printer
type Greeter struct {
	Printer Printer
}

// Define a method for our greeter
func (g Greeter) Greet() {
	g.Printer.Print()
}

func main() {
	//Create a fancy message printer
	fancyPrinter := FancyPrinter{Message: "Hello, world"}

	// Create a greeter and inject the fancy message printer
	greeter := Greeter{Printer: &fancyPrinter}

	//Call the Greet method on the greeter
	greeter.Greet()

	msgPrinter := MessagePrinter{Message: "Hello, Message Printer"}

	greeter = Greeter{Printer: &msgPrinter}
	greeter.Greet()
}
