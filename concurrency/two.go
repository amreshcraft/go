package main

import (
	"fmt"
	"net/http"
	"runtime"
	"time"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Request started:", time.Now())
	time.Sleep(2 * time.Second)
	fmt.Fprintf(w, "Request completed at: %v\n", time.Now())
}

func main() {
	runtime.GOMAXPROCS(8) // Force Go to use only 1 OS thread

	http.HandleFunc("/", handler)

	fmt.Println("Starting server on :8080...")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Server error:", err)
	}
}
