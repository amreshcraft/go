package main

import "fmt"

type Shape interface {
	Area() float64
}

type Rectangle struct{ Width, Height float64 }
type Circle struct{ Radius float64 }

func (r Rectangle) Area() float64 { return r.Width * r.Height }
func (c Circle) Area() float64    { return 3.14 * c.Radius * c.Radius }

func printArea(s Shape) {
	fmt.Println("Area:", s.Area())
}

func main() {
	r := Rectangle{10, 5}
	c := Circle{7}

	printArea(r)
	printArea(c)
}
