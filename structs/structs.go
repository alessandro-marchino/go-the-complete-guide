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

func (u user) outputUserDetails() {
	fmt.Println(u.firstName, u.lastName, u.birthdate)
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

	// Do something awesome with that gathered data!
	appUser.outputUserDetails()
}

func getUserData(promptText string) (value string) {
	fmt.Print(promptText)
	fmt.Scan(&value)
	return value
}
