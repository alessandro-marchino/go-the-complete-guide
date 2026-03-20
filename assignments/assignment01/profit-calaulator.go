package main

import "fmt"

func main() {
	var revenue float64
	fmt.Print("Enter revenue: ")
	fmt.Scan(&revenue)

	var expenses float64
	fmt.Print("Enter expenses: ")
	fmt.Scan(&expenses)

	var taxRate float64
	fmt.Print("Enter tar rate: ")
	fmt.Scan(&taxRate)

	earningsBeforeTax := revenue - expenses
	earningsAfterTax := earningsBeforeTax * (1 - taxRate / 100)
	ratio := earningsBeforeTax / earningsAfterTax

	fmt.Print("Earnings Before Tax: ")
	fmt.Println(earningsBeforeTax)

	fmt.Print("Earnings After Tax: ")
	fmt.Println(earningsAfterTax)

	fmt.Print("Ratio: ")
	fmt.Println(ratio)
}