package main

import "fmt"

func init(){
	fmt.Println("Hello init func 1")
}

func init(){
	fmt.Println("Hello init func 2")
}
func init(){
	fmt.Println("Hello init func 3")
}
func init(){
	fmt.Println("Hello init func 4")
}
func init(){
	fmt.Println("Hello init func 5")
}
func init(){
	fmt.Println("Hello init func 6")
}
func main() {
	fmt.Println("Inside the main func")
}