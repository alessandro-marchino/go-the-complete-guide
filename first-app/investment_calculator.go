package main

import (
	"fmt"
	"math"
)

const INFLATION = 2.5

func main() {
	var investmentAmount float64
	var years float64
	var expectedReturnRate float64

	outputText("Investment amount: ")
	fmt.Scan(&investmentAmount)

	outputText("Expected return rate: ")
	fmt.Scan(&expectedReturnRate)

	outputText("Inverment duration in years: ")
	fmt.Scan(&years)

	futureValue, futureRealValue := calculateFutureValues(investmentAmount, expectedReturnRate, years)

	fmt.Printf("Future value: %.2f\nFuture value (adjusted for inflation): %.2f\n", futureValue, futureRealValue)
}

func outputText(text string) {
	fmt.Print(text)
}

func calculateFutureValues(investmentAmount float64, expectedReturnRate float64, years float64) (fv float64, rfv float64) {
	fv = investmentAmount * math.Pow(1 + expectedReturnRate / 100, years)
	rfv = fv / math.Pow(1 + INFLATION / 100, years)
	return fv, rfv
}
