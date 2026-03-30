package main

import (
	"fmt"

	"example.com/structs/user"
)

type customString string
func (text *customString) log() {
	fmt.Println("[LOG]", *text)
}

func main() {
	var name customString
	name = "Max"
	name.log()

	userFirstName := getUserData("Please enter your first name: ")
	userLastName := getUserData("Please enter your last name: ")
	userBirthdate := getUserData("Please enter your birthdate (DD/MM/YYYY): ")

	appUser, err := user.New(userFirstName, userLastName, userBirthdate)

	if err != nil {
		fmt.Println(err)
		return
	}
	// Do something awesome with that gathered data!
	appUser.OutputUserDetails()

	appUser.ClearUserName()
	appUser.OutputUserDetails()

	admin := user.NewAdmin("test@example.com", "password")
	admin.OutputUserDetails()

}

func getUserData(promptText string) (value string) {
	fmt.Print(promptText)
	fmt.Scanln(&value)
	return value
}
