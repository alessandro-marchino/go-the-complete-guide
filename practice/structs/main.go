package main

import (
	"fmt"

	"example.com/practice/structs/note"
)

func main() {
	title, content := getNodeData()
	note, err := note.New(title, content)
	if err != nil {
		fmt.Println(err)
		return
	}
}

func getNodeData() (title, content string) {
	title = getUserInput("Note title:")
	content = getUserInput("Note content:")
	return title, content
}

func getUserInput(prompt string) (value string) {
	fmt.Print(prompt)
	fmt.Scanln(&value)
	return value
}
