package main

import (
	"fmt"
	"time"
)

func greet(phrase string) {
	fmt.Println("Hello!", phrase)
}

func slowGreet(phrase string) {
	// Simulate a slow, long-taking task
	time.Sleep(3 * time.Second)
	fmt.Println("Hello!", phrase)
}

func main() {
	go greet("Nice to meet you!")
	go greet("How are you?")
	go slowGreet("How ... are ... you ...?")
	go greet("I hope you're liking the course")
}
