package main

import (
	"errors"
	"fmt"
	"strings"
)

func someFunction() (string, error) {
	return "", errors.New("specific error message")
}

func somethingGoesWrong() bool {
	return true
}

func someFunctionReturnError() error {
	if somethingGoesWrong() {
		return fmt.Errorf("something went wrong")
	}
	return nil
}

func someItemFunction(item int) error {
	return errors.New(fmt.Sprintf("Item error %d", item))
}

type MyError struct {
	Message string
	Code    int
}

func (e *MyError) Error() string {
	return fmt.Sprintf("error: %s (code %d)", e.Message, e.Code)
}

func someCustomErrorfunction() error {
	if somethingGoesWrong() {
		return &MyError{Message: "something went wrong", Code: 123}
	}
	return nil
}

func main() {
	// Checking for specific error
	_, err := someFunction()
	if err != nil {
		if strings.Contains(err.Error(), "specific error message") {
			fmt.Println("handle specific error")
		} else {
			fmt.Println("handle other errors")
		}
	}

	// Returning errors
	err = someFunctionReturnError()
	if err != nil {
		fmt.Println(err.Error())
	}

	// handling multiple error
	items := []int{1, 2, 3, 4, 5}
	errs := []error{}
	for _, item := range items {
		err := someItemFunction(item)
		if err != nil {
			errs = append(errs, err)
		}
	}
	if len(errs) > 0 {
		fmt.Println(errs)
	}

	//Defining custom errors
	err = someCustomErrorfunction()
	if err != nil {
		if myErr, ok := err.(*MyError); ok {
			fmt.Println(myErr.Error())
		} else {
			fmt.Println(myErr)
		}
	}
}
