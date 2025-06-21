package main

import (
	"fmt"
	"os"
)

func main() {
	testExit()
}

func testExit(){
	defer fmt.Println("Hello defer") // not executed
	os.Exit(1)
	panic("after the os.exit") // not executed
}