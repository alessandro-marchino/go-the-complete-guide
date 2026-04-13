package filemanager

import (
	"bufio"
	"errors"
	"os"
)

func ReadLines(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, errors.New("Failed to open file")
	}
	scanner := bufio.NewScanner(file)
	var lines []string
	for scanner.Scan() {
		text := scanner.Text()
		if(len(text) > 0) {
			lines = append(lines, scanner.Text())
		}
	}
	file.Close()
	err = scanner.Err()
	if err != nil {
		return nil, errors.New("Failed to read line in file")
	}
	return lines, nil
}
