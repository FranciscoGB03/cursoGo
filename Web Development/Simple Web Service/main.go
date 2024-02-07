package main

import (
	"fmt"
	"net/http"
)

func Handle(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, World!")
}

func main() {
	http.HandleFunc("/", Handle)

	http.ListenAndServe(":8000", nil)
}
