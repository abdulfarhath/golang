package main

import (
	"fmt"
	"net/http"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, World!")
}

func main() {
	http.HandleFunc("/", helloHandler) // Set route and handler

	fmt.Println("Server starting at http://localhost:8080")

	err := http.ListenAndServe(":8080", nil) // Start server on port 8080
	if err != nil {
		fmt.Println("Error starting server:", err)
	}
}