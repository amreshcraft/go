package main

import (
	"fmt"
	"time"
)

func worker(id int) {
	for i := 1; i <= 5; i++ {
		fmt.Printf("Worker %d: step %d\n", id, i)
		time.Sleep(500 * time.Millisecond) // Simulate long work
	}
}

func main() {
	fmt.Println("Starting workers...")

	go worker(1)
	go worker(2)
	go worker(3)

	// Wait for all workers to finish (simplest way for now)
	time.Sleep(5 * time.Second)
	fmt.Println("All workers done.")
}
