package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {
	// Make a GET request
	resp, err := http.Get("https://api.github.com")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer resp.Body.Close()

	// Read and print the response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading body:", err)
		return
	}
	fmt.Println(string(body))
}