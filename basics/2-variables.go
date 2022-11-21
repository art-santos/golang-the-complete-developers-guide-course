package main

import "fmt"

func main() {
	var favoriteColor string = "red"
	birthYear, age := 2000, 22
	var (
		initial string = "a"
		final string = "r"
	)
	var ageInDays int

	ageInDays = age*365

	fmt.Println("favorite color:", favoriteColor, "birthday:", birthYear, "age:", age, "initial:", initial, "final:", final, "age in days:", ageInDays)
}