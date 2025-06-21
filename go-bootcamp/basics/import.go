package main

import (
	print "fmt"
	"net/http"
)

func ImportTest() {
	print.Println("Hello")
	res,err := http.Get("https://amreshmaurya.com")
	if err != nil{
		panic("Error while fetching")
	}
	defer res.Body.Close();
	print.Println("Status Code : " , res.StatusCode)
	print.Println("Data: ",res)
}