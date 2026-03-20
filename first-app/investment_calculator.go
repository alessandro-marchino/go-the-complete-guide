package main

import (
	"fmt"
	"math"
)

func main() {
	const INFLATION = 2.5
	var investmentAmount float64
	var years float64
	var expectedReturnRate float64

	outputText("Investment amount: ")
	fmt.Scan(&investmentAmount)

	outputText("Expected return rate: ")
	fmt.Scan(&expectedReturnRate)

	outputText("Inverment duration in years: ")
	fmt.Scan(&years)

	futureValue := investmentAmount * math.Pow(1.0 + expectedReturnRate / 100, years)
	futureRealValue := futureValue / math.Pow(1 + INFLATION / 100, years)

	fmt.Printf("Future value: %.2f\nFuture value (adjusted for inflation): %.2f\n", futureValue, futureRealValue)
}

func outputText(text string) {
	fmt.Print(text)
}
