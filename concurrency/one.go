package main

import (
	"fmt"
	"net/http"
	"time"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Request received at:", time.Now())
	time.Sleep(2 * time.Second) // Simulate long processing
	fmt.Fprintf(w, "Hello from Goroutine!")
}

func main() {
	http.HandleFunc("/", handler)

	fmt.Println("Starting server on :8080...")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Server error:", err)
	}
}
