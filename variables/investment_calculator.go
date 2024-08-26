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

	// fmt.Print("Enter your investment amount: ")
	outputText("Investment amount: ")
	fmt.Scan(&investmentAmount)

	// fmt.Print("Years: ")
	outputText("Years: ")
	fmt.Scan(&years)

	// fmt.Print("Expected Return Rate: ")
	outputText("Expected Return Rate: ")
	fmt.Scan(&expectedReturnRate)

	futureValue := investmentAmount * math.Pow(1 + expectedReturnRate / 100, years)

	futureRealValue := futureValue / math.Pow((1 + inflationRate / 100), years)

	// formattedFV := fmt.Sprintf("Future value: %.0f\n", futureValue)
	// formattedRFV := fmt.Sprintf("Future value (adjusted for inflation): %.0f\n", futureRealValue)
	// fmt.Println("Future value:", futureValue)
	fmt.Printf("Future value: %.0f\nFuture value (adjusted for inflation): %.0f\n", futureValue, futureRealValue)
	// fmt.Println("Future value (adjusted for inflation):", futureRealValue)
	// fmt.Print(formattedFV, formattedRFV)
}

func outputText(text string) {
	fmt.Print(text)
}