package main

import (
	"fmt"
	"math"
	"net/http"
	"runtime"
	"time"
)

func heavyTask(id int) {
	fmt.Printf("Request %d started at: %v\n", id, time.Now())
	result := 0.0
	// Simulate heavy CPU work
	for i := 0; i < 1e9; i++ {
		result += math.Sqrt(float64(i))
	}
	fmt.Printf("Request %d finished at: %v\n", id, time.Now())
}

func handler(w http.ResponseWriter, r *http.Request) {
	heavyTask(1)
	fmt.Fprintf(w, "Request processed at: %v\n", time.Now())
}

func main() {
	// Try changing this to 1, 2, 4 to see the difference
	runtime.GOMAXPROCS(4)

	http.HandleFunc("/", handler)

	fmt.Println("Starting server on :8080...")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Server error:", err)
	}
}
