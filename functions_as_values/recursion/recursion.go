package recursion

import "fmt"

func main() {
	fact := factorial(3);

	fmt.Println(fact)
}

func factorial(number int) int {
	if number == 0 {
		return number
	}
	return number * factorial(number - 1)
}