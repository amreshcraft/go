package main

import (
	"fmt"
	"reflect"
)


func main() {
// map[keytype]valuetype
	m := make(map[int]string);
	m[1] = "Amresh";
	m[2] = "Ojas";
	m[3] = "Aniket";

	fmt.Println(reflect.TypeOf(m))
	fmt.Println(m[1])
	delete(m,2);
	fmt.Println(m)
	clear(m)
	fmt.Println(m)
	
}