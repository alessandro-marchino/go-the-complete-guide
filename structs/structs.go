package main

import (
	"fmt"
	"time"
)

type user struct {
	firstName string
	lastName string
	birthdate string
	createdAt time.Time
}

func main() {
	userFirstName := getUserData("Please enter your first name: ")
	userLastName := getUserData("Please enter your last name: ")
	userBirthdate := getUserData("Please enter your birthdate (DD/MM/YYYY): ")

	var appUser user
	appUser = user{
		firstName: userFirstName,
		lastName: userLastName,
		birthdate: userBirthdate,
		createdAt: time.Now(),
	}

	// Do something awasome with that gathered data!
	outputUserDetails(&appUser)
}

func outputUserDetails(appUser *user) {
	fmt.Println(appUser.firstName, appUser.lastName, appUser.birthdate)
}

func getUserData(promptText string) (value string) {
	fmt.Print(promptText)
	fmt.Scan(&value)
	return value
}
