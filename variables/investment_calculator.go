package main

import (
	"fmt"
	"math"
)

func main() {
	const inflationRate = 2.5
	var investmentAmount float64
	var years float64 = 10.0
	var expectedReturnRate float64 = 5.5

	fmt.Print("Enter your investment amount: ")
	fmt.Scan(&investmentAmount)

	fmt.Print("Years: ")
	fmt.Scan(&years)

	fmt.Print("Expected Return Rate: ")
	fmt.Scan(&expectedReturnRate)

	futureValue := investmentAmount * math.Pow(1 + expectedReturnRate / 100, years)

	futureRealValue := futureValue / math.Pow((1 + inflationRate / 100), years)

	// fmt.Println("Future value:", futureValue)
	fmt.Printf("Future value: %.0f\nFuture value (adjusted for inflation): %.0f\n", futureValue, futureRealValue)
	// fmt.Println("Future value (adjusted for inflation):", futureRealValue)
}