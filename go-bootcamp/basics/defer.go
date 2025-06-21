package main

import "fmt"

func main() {
input(10)

// note : first defer evaluted first so on the evaluation time i = 10 , ye wahi value print krega jo evaluation k time p thi
// output
// Before update: value of i :  10
// After Update : value of i :  11
// Third defer : value of i :  10
// Second defer : value of i :  10
// First defer : value of i :  10
}

func input(i int) {
	defer fmt.Println("First defer : value of i : ", i) // 10
	defer fmt.Println("Second defer : value of i : ", i) // 10
	defer fmt.Println("Third defer : value of i : ", i) // 10
	fmt.Println("Before update: value of i : ", i)
	i=i+1;
	fmt.Println("After Update : value of i : ", i)

}
