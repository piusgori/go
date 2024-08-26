package main

import (
	"fmt"
	"math"
)

const inflationRate = 2.5

func main() {
	var investmentAmount float64
	var years float64 = 10.0
	expectedReturnRate := 5.5
	
	// fmt.Print("Enter your investment amount: ")
	outputText("Investment amount: ")
	fmt.Scan(&investmentAmount)

	// fmt.Print("Years: ")
	outputText("Years: ")
	fmt.Scan(&years)

	// fmt.Print("Expected Return Rate: ")
	outputText("Expected Return Rate: ")
	fmt.Scan(&expectedReturnRate)

	futureValue, futureRealValue := calculateFutureValues(investmentAmount, expectedReturnRate, years)

	// futureValue := investmentAmount * math.Pow(1 + expectedReturnRate / 100, years)

	// futureRealValue := futureValue / math.Pow((1 + inflationRate / 100), years)

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

func calculateFutureValues(investmentAmount, expectedReturnRate, years float64) (float64, float64) {
	fv := investmentAmount * math.Pow(1 + expectedReturnRate / 1000, years)
	rfv := fv / math.Pow((1 + inflationRate / 100), years)
	return fv, rfv
}