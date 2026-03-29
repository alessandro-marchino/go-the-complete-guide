package main

import "fmt"

func main() {
	firstName := getUserData("Please enter your first name: ")
	lastName := getUserData("Please enter your last name: ")
	birthdate := getUserData("Please enter your birthdate (DD/MM/YYYY): ")

	// Do something awasome with that gathered data!
	fmt.Println(firstName, lastName, birthdate)
}

func getUserData(promptText string) (value string) {
	fmt.Print(promptText)
	fmt.Scan(&value)
	return value
}
