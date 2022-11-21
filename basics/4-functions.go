package main 

import "fmt"

func double(x int) int {
	return x + x
}

func add(lhs int, rhs int) int {
	return lhs + rhs
}

func addThree(a, b, c int) int {
	return a + b + c
}

func greetPerson(name string) {
	fmt.Println(name)
}

func hiThere() string {
	return "Hi there!"
}

func main(){
	fmt.Println(add(1,2), double(1), addThree(1,2,3), hiThere())
	greetPerson("under")
}