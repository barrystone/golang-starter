package main

import "fmt"

type bot interface {
	getGreeting() string
	// getGreeting(string, int) (string, error)
	// getBotVersion() float64
	// respondToUser(userInput string) string
}

type englishBot struct{}
type spanishBot struct{}

func main() {
	eb := englishBot{}
	sb := spanishBot{}

	printGreeting(eb)
	printGreeting(sb)
}

func printGreeting(b bot) {
	fmt.Println(b.getGreeting())
}

// func (eb englishBot) getGreeting() string {
func (englishBot) getGreeting() string {
	return "Hi there!"
}
func (spanishBot) getGreeting() string {
	return "Hola!"
}
