package main

type Person struct {
	name string
	age  int
}

type Employee struct {
	Person
	salary float64
}

type EmployeePtr struct {
	*Person
	salary float64
}

type Employee_Anonymous struct {
	Person
	salary float64
	string // anonymous field
}

type Writer interface {
	Write([]byte) (int, error)
}

type Closer interface {
	Close() error
}

type WriteCloser interface {
	Writer
	Closer
}
