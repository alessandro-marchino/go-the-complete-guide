package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"example.com/practice/structs/note"
	"example.com/practice/structs/todo"
)

type saver interface {
	Save() error
}

func main() {
	title, content := getNodeData()
	note, err := note.New(title, content)
	if err != nil {
		fmt.Println(err)
		return
	}
	note.Display()

	err = note.Save()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Saving the note succeeded")

	todoText := getTodoData()
	todo, err := todo.New(todoText)
	if err != nil {
		fmt.Println(err)
		return
	}
	todo.Display()
	err = todo.Save()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Saving the todo succeeded")
}

func getNodeData() (title, content string) {
	title = getUserInput("Note title:")
	content = getUserInput("Note content:")
	return title, content
}

func getTodoData() string {
	return getUserInput("Todo text:")
}

func getUserInput(prompt string) (value string) {
	fmt.Printf("%v ", prompt)

	reader := bufio.NewReader(os.Stdin)
	text, err := reader.ReadString('\n')

	if err != nil {
		return ""
	}
	text = strings.TrimSuffix(text, "\n")
	value = strings.TrimSuffix(text, "\r")
	return value
}
