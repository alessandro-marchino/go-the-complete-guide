package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

const ACCOUNT_BALANCE_FILE = "balance.txt"

func writeFloatToFile(value float64, fileName string) {
	balanceText := fmt.Sprint(value)
	os.WriteFile(fileName, []byte(balanceText), 0644)
}

func getFloatFromFile(fileName string) (float64, error) {
	data, err := os.ReadFile(fileName)
	if err != nil {
		return 1000, errors.New("Failed to find file")
	}
	valueText := string(data)
	value, err := strconv.ParseFloat(valueText, 64)
	if err != nil {
		return 1000, errors.New("Failed to parse stored value")
	}
	return value, nil
}

func main() {
	var accountBalance, err = getFloatFromFile(ACCOUNT_BALANCE_FILE)
	if err != nil {
		fmt.Println("ERROR!")
		fmt.Println(err)
		fmt.Println("----------")
	}

	fmt.Println("Welcome to Go Bank!")

	for {
		presentOptions()

		var choice int
		fmt.Print("Your choice: ")
		fmt.Scan(&choice)

		switch choice {
		case 1:
			fmt.Printf("Your balance is %.2f\n", accountBalance)
		case 2:
			fmt.Print("Your deposit: ")
			var depositAmount float64
			fmt.Scan(&depositAmount)

			if depositAmount <= 0 {
				fmt.Println("Invalid amount. Must be greater than 0")
				continue
			}
			accountBalance += depositAmount
			fmt.Printf("Balance updated|\nour balance is %.2f\n", accountBalance)
			writeFloatToFile(accountBalance, ACCOUNT_BALANCE_FILE)
		case 3:
			fmt.Print("Withdraw amount: ")
			var withdrawnAmount float64
			fmt.Scan(&withdrawnAmount)
			if withdrawnAmount <= 0 {
				fmt.Println("Invalid amount. Must be greater than 0")
				continue
			}
			if withdrawnAmount > accountBalance {
				fmt.Println("Invalid amount. You can't withdraw more than you have")
				continue
			}
			accountBalance -= withdrawnAmount
			fmt.Printf("Balance updated|\nour balance is %.2f\n", accountBalance)
			writeFloatToFile(accountBalance, ACCOUNT_BALANCE_FILE)
		default:
			fmt.Println("Goodbye!")
			fmt.Println("Thanks for choosing our bank")
			return
		}
	}
}

