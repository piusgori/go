package functionsarevalues

import "fmt"

type transformFn func(int) int

// type anotherFn func(int, []string, map[string][]int) ([]int, string)

func main() {
	numbers := []int{1,2,3,4}
	moreNumbers := []int{5, 1, 2}

	doubled := transformNumbers(&numbers, double);
	trippled := transformNumbers(&numbers, tripple);

	fmt.Println(doubled, trippled)

	transformerFn1 := getTransformerFunction(&numbers)
	transformerFn2 := getTransformerFunction(&moreNumbers)


	transformedNumbers := transformNumbers(&numbers, transformerFn1)
	moreTransformedNumbers := transformNumbers(&moreNumbers, transformerFn2)

	fmt.Println(transformedNumbers, moreTransformedNumbers)
}

func transformNumbers(numbers *[]int, transform transformFn) []int {
	dNumbers := []int{}
	for _, val := range *numbers {
		dNumbers = append(dNumbers, transform(val))
	}

	return dNumbers
}

func getTransformerFunction (numbers *[]int) transformFn {
	if (*numbers)[0] == 1 {
		return double
	} else {
		return tripple
	}
}

func double(number int) int {
	return number * 2
}

func tripple(number int) int {
	return number * 3
}