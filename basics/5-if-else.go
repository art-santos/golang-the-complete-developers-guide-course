package main

import "fmt"

const (
	Monday = 0
	Tuesday = 1
	Wednesday = 2
	Thursday = 3
	Friday = 4
	Saturday = 5
	Sunday = 6
)

const (
	Admin = 10
	Manager = 20
	Contractor = 30
	Member = 40
	Guest = 50
)

func weekday(day int) bool{
	return day <= 4
}

func accessDenied() string {
	return "denied"
}

func accessGranted() string {
	return "granted"
}

func isAllowed(today int, role int) string {
	if(weekday(today)){
		return accessGranted()
	}else if (!weekday(today) && role < 30){
		return accessGranted()
	}
		return accessDenied()
}

func main(){
	fmt.Println(isAllowed(Sunday, Member))
}