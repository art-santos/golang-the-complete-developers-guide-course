package main

import "fmt"


func multiReturn() (int, int, int){
	return 1,2,3
}

func main() {
	a,b, _ := multiReturn()

	fmt.Println(a,b)
}