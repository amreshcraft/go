package main

import "fmt"

func main() {
	for i := 0; i < 5; i++ {
	go func(i int){
		fmt.Println("My id : " , i)
	}(i)

	
	}
}