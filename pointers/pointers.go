package main

import "fmt"

func main() {
	age := 32

	var agePointer *int = &age

	fmt.Println("Age:", agePointer)
	adultYears := getAdultYears(age)
	fmt.Println(adultYears)
}

func getAdultYears(age int) int {
	return age - 18
}