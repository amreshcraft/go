package main

import "fmt"

type Animal struct {
	Name string
}

func (a Animal) Speak() {
	fmt.Println(a.Name, "makes a sound")
}

type Dog struct {
	Animal // Embedded struct (Composition)
}

func (d Dog) Bark() {
	fmt.Println(d.Name, "barks")
}

func main() {
	d := Dog{Animal{Name: "Tommy"}}
	d.Speak() // Inherited-like behavior
	d.Bark()
}
