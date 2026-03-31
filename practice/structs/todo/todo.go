package todo

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
)

type Todo struct {
	Text string `json:"text"`
}

func New(text string) (*Todo, error) {
	if text == "" {
		return nil, errors.New("Invalid input")
	}
	return &Todo {
		Text: text,
	}, nil
}

func (t *Todo) Display() {
	fmt.Printf(t.Text)
}

func (t *Todo) Save() error {
	fileName := "todo.json"
	jsonData, err := json.Marshal(*t)
	if err != nil {
		return err
	}
	return os.WriteFile(fileName, jsonData, 0644)
}
