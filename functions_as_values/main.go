package main

import "fmt"

type transformFn func(int) int

// type anotherFn func(int, []string, map[string][]int) ([]int, string)

func main() {
	numbers := []int{1,2,3,4}

	doubled := transformNumbers(&numbers, double);

	trippled := transformNumbers(&numbers, tripple);

	fmt.Println(doubled, trippled)
}

func transformNumbers(numbers *[]int, transform transformFn) []int {
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