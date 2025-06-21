package main

import "fmt"

func sum(nums... int )int{
	res:=0
	for i:=  range nums {
		res += i
	}
	return res
}

func main() {
	fmt.Println(sum(1,2,3,4,5,6,7,8,9))
}