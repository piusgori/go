package main

import "fmt"

func main() {
	numbers := []int{1,2,3,4}

	doubled := transformNumbers(&numbers, double);

	trippled := transformNumbers(&numbers, tripple);

	fmt.Println(doubled, trippled)
}

func transformNumbers(numbers *[]int, transform func(int) int) []int {
	dNumbers := []int{}
	for _, val := range *numbers {
		dNumbers = append(dNumbers, transform(val))
	}

	return dNumbers
}

func double(number int) int {
	return number * 2
}

func tripple(number int) int {
	return number * 3
}