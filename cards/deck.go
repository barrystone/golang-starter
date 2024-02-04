package main

import (
	"fmt"
	"os"
	"strings"
)

// Creat a new type of 'deck'
// which is a slice of strings
type deck []string

func newDeck() deck {
	cards := deck{}

	cardSuits := []string{"Spades", "Diamods", "Hearts", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three", "Four"}

	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, value+" of "+suit)
		}
	}

	return cards
}

// Receiver function
// 'd' is a variable(like 'this' or 'self'),
// excetly like cards put in there
func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

func (d deck) toString() string {
	// https://pkg.go.dev/strings#example-Join
	return strings.Join([]string(d), ",")
}

func (d deck) saveToFile(fileName string) error {
	// https://pkg.go.dev/os#WriteFile
	return os.WriteFile(fileName, []byte(d.toString()), 0666)
	// "0666" means anyone can read and write this file
}

func newDeckFromFile(fileName string) deck {
	// https://pkg.go.dev/os#ReadFile
	bs, err := os.ReadFile(fileName)
	if err != nil {
		// Option #1 - log the error and return a call to newDeck()
		// Option #2 - log the error and entirely quit the program
		fmt.Println("Error: ", err)
		os.Exit(1)
	}
	//string(bs): Ace of Spades,Two of Spades,Three of Spades, ...
	// https://pkg.go.dev/strings@go1.21.6#Split
	s := strings.Split(string(bs), ",")
	return deck(s)
}
