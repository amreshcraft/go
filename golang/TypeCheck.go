package main
import (
"fmt"
"reflect"
)

func main() {
	var n int  = 10;
	var f float64  = 23.4;
	var isAdmin bool = false;
	var m map[string] int ;
	var bte byte = 3
	
	fmt.Println(reflect.TypeOf(n))
	fmt.Println(reflect.TypeOf(f))
	fmt.Println(reflect.TypeOf(isAdmin))
	fmt.Println(reflect.TypeOf(m))
	fmt.Println(reflect.TypeOf(bte))

}