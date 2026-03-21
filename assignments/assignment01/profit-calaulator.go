package main

import "fmt"

func main() {
	revenue := getUserInput("Enter revenue: ")
	expenses := getUserInput("Enter expenses: ")
	taxRate := getUserInput("Enter tar rate: ")

	earningsBeforeTax, earningsAfterTax, ratio := calculateFinancials(revenue, expenses, taxRate)

	fmt.Printf("Earnings Before Tax: %.2f\n", earningsBeforeTax)
	fmt.Printf("Earnings After Tax: %.2f\n", earningsAfterTax)
	fmt.Printf("Ratio: %.3f\n", ratio)
}

func getUserInput(message string) (res float64) {
	fmt.Print(message)
	fmt.Scan(&res)
	return res
}

func calculateFinancials(revenue, expenses, taxRate float64) (ebt, eat, ratio float64) {
	ebt = revenue - expenses
	eat = ebt * (1 - taxRate / 100)
	ratio = ebt / eat
	return ebt, eat, ratio
}
