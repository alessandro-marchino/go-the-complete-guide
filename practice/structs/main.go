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
type outputtable interface {
	Display()
	saver
}

func main() {
	title, content := getNodeData()
	note, err := note.New(title, content)
	if err != nil {
		fmt.Println(err)
		return
	}
	err = outputData(note)
	if err != nil {
		return
	}

	todoText := getTodoData()
	todo, err := todo.New(todoText)
	if err != nil {
		fmt.Println(err)
		return
	}
	err = outputData(todo)
	if err != nil {
		return
	}
}

func outputData(data outputtable) error {
	data.Display()
	return saveData(data)
}

func saveData(data saver) error {
	err := data.Save()
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}
func printSomething(value interface{}) {
	fmt.Println(value)
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
