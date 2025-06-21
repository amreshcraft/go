package main

import "fmt"

func main() {
	testapi()
	fmt.Println("My other import task ")
}

func testapi(){
	defer func(){
     if i:= recover(); i != nil{
		fmt.Println("my recovered msg:" ,i)
	 }
	}()

	panic("Kuch to gadbad h")
	fmt.Println("unreachable code")
}