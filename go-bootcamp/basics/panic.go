package main

import "fmt"

func main() {
	input(5)
	input(-3)
}

func input(i int) {
	defer fmt.Println("Defer 1")
	defer fmt.Println("defer 2")
	fmt.Println("Given value : ", i)
	if i < 0 {
		panic("Negative number are not allowed")
	}
}
