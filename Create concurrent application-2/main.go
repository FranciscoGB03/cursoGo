package main

import (
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	//Create a new channel to receive the response from a goroutine
	responseChan := make(chan string)

	// Start a new goroutine to process the request
	go processRequest(r, responseChan)

	// Wait for the response from the goroutine
	response := <-responseChan

	//Write the response to the client
	fmt.Fprintf(w, response)
}

func processRequest(r *http.Request, responseChan chan string) {
	// Process the request and generate a response
	response := "Hello, Golang World!"

	// Send the response back to the main goroutine via the channel
	responseChan <- response
}

func main() {
	// Register the handler function to handle requests for the root URL path "/"
	http.HandleFunc("/", handler)

	// Start the server listening on port 8080
	err := http.ListenAndServe(":8081", nil)
	if err != nil {
		panic(err)
	}
}
