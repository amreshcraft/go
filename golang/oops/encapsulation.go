

package main

import "fmt"

type Person struct {
	Name string // Public
	age  int    // Private
}

func main() {
	p := Person{Name: "Amresh", age: 24}
	fmt.Println(p.Name) // Accessible
	fmt.Println(p.age) // Not accessible from other packages
}
