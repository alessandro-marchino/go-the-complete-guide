package note

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"strings"
	"time"
)

type Note struct {
	title string
	content string
	createdAt time.Time
}

func New(title, content string) (*Note, error) {
	if title == "" || content == "" {
		return nil, errors.New("Invalid input")
	}
	return &Note {
		title: title,
		content: content,
		createdAt: time.Now(),
	}, nil
}

func (n *Note) Display() {
	fmt.Printf("Your note titled \"%v\" has the following content:\n\n%v\n", n.title, n.content)
}

func (n *Note) Save() error {
	fileName := strings.ReplaceAll(n.title, " ", "_")
	fileName = strings.ToLower(fileName)
	jsonData, err := json.Marshal(*n)
	if err != nil {
		return err
	}
	return os.WriteFile(fileName, jsonData, 0644)
}
