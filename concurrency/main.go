package main

import (
	"fmt"
	"time"
)

func greet(phrase string) {
	fmt.Println("Hello!", phrase)
}

func slowGreet(phrase string, doneChan chan bool) {
	// Simulate a slow, long-taking task
	time.Sleep(3 * time.Second)
	fmt.Println("Hello!", phrase)

	doneChan <- true
}

func main() {
	go greet("Nice to meet you!")
	go greet("How are you?")
	done := make(chan bool)
	go slowGreet("How ... are ... you ...?", done)
	go greet("I hope you're liking the course")
	<- done
}
