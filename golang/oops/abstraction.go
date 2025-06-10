package main

import "fmt"

type Shape interface {
	Area() float64
}

type Rectangle struct {
	Width  float64
	Height float64
}

func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

func main() {
	var s Shape = Rectangle{Width: 10, Height: 5}
	fmt.Println("Area:", s.Area())
}
