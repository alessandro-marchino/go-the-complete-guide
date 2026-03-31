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
	Title string
	Content string
	CreatedAt time.Time
}

func New(title, content string) (*Note, error) {
	if title == "" || content == "" {
		return nil, errors.New("Invalid input")
	}
	return &Note {
		Title: title,
		Content: content,
		CreatedAt: time.Now(),
	}, nil
}

func (n *Note) Display() {
	fmt.Printf("Your note titled \"%v\" has the following content:\n\n%v\n", n.Title, n.Content)
}

func (n *Note) Save() error {
	fileName := strings.ReplaceAll(n.Title, " ", "_")
	fileName = strings.ToLower(fileName) + ".json"
	jsonData, err := json.Marshal(*n)
	if err != nil {
		return err
	}
	return os.WriteFile(fileName, jsonData, 0644)
}
