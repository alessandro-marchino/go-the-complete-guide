package main

import (
	"errors"
	"fmt"
	"os"
)

func main() {
	revenue, err := getUserInput("Enter revenue: ")
	if err != nil {
		fmt.Println(err)
		return
	}
	expenses, err := getUserInput("Enter expenses: ")
	if err != nil {
		fmt.Println(err)
		return
	}
	taxRate, err := getUserInput("Enter tar rate: ")
	if err != nil {
		fmt.Println(err)
		return
	}

	// Goals:
	// 1) Validate user input
	// 	=> Show error messages & exit if invalid input is provided
	// 	- no negative numbers
	// 	- no 0
	// 2) Store calculated results into file

	earningsBeforeTax, earningsAfterTax, ratio := calculateFinancials(revenue, expenses, taxRate)

	fmt.Printf("Earnings Before Tax: %.2f\n", earningsBeforeTax)
	fmt.Printf("Earnings After Tax: %.2f\n", earningsAfterTax)
	fmt.Printf("Ratio: %.3f\n", ratio)

	writeToFile(revenue, expenses, taxRate, earningsBeforeTax, earningsAfterTax, ratio)
}

func getUserInput(message string) (res float64, err error) {
	fmt.Print(message)
	fmt.Scan(&res)
	if res <= 0 {
		return 0, errors.New("Invalid input")
	}
	return res, nil
}

func calculateFinancials(revenue, expenses, taxRate float64) (ebt, eat, ratio float64) {
	ebt = revenue - expenses
	eat = ebt * (1 - taxRate / 100)
	ratio = ebt / eat
	return ebt, eat, ratio
}

func writeToFile(revenue, expenses, taxRate, ebt, eat, ratio float64) (err error) {
	message := fmt.Sprintf("Revenue: %.2f\nExpenses: %.2f\nTax Rate: %.2f\nEarnings Before Tax: %.2f\nEarnings After Tax: %.2f\nRatio: %.3f\n",
		revenue, expenses, taxRate, ebt, eat, ratio)
	messageBytes := []byte(message)
	return os.WriteFile("result.txt", messageBytes, 0644)
}
