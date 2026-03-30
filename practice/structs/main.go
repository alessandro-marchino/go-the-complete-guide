package main

import (
	"errors"
	"fmt"
)

func main() {
	title, content, err := getNodeData()
	if err != nil {
		fmt.Println(err)
		return
	}
}

func getNodeData() (title string, content string, err error) {
	title, err = getUserInput("Note title:")
	if err != nil {
		return "", "", err
	}
	content, err = getUserInput("Note content:")
	if err != nil {
		return "", "", err
	}
	return title, content, nil
}

func getUserInput(prompt string) (value string, err error) {
	fmt.Print(prompt)
	fmt.Scanln(&value)

	if value == "" {
		return "", errors.New("Invalid input")
	}
	return value, nil
}
