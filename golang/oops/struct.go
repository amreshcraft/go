package main

import "fmt"

type Person struct {
	name   string
	age    int
	salary float64
}

func main() {
	person := Person{
		name:   "Amresh",
		age:    24,
		salary: 20000,
	}
	fmt.Println(person) // {Amresh 24 20000} // person is object
}
