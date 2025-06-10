package main

import "fmt"

type Shape interface {
	Area() float64
	Perimeter() float64
}

type Circle struct {
	Radius int
}

func (c Circle) Perimeter() float64 {
	return 2 * 3.14 * float64(c.Radius)
}

func (c Circle) Area() float64 {
	return 3.14 * float64(c.Radius) * float64(c.Radius)
}

func main() {
	var s Shape = Circle{
		Radius: 5,
	}
	fmt.Println("Area: ", s.Area())
	fmt.Println("Perimeter: ", s.Perimeter())
}
