package main

import (
	"fmt"
	"strings"
)

// Creat a new type of 'deck'
// which is a slice of strings
type deck []string

func newDeck() deck {
	cards := deck{}

	cardSuits := []string{"Spades","Diamods","Hearts","Clubs"}
	cardValues := []string{"Ace","Two","Three","Four"}

	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, value + " of " + suit)
		}
	}

	return cards
}

// Receiver function
// 'd' is a variable(like 'this' or 'self'),
// excetly like cards put in there
func (d deck) print(){
	for i, card := range d {
		fmt.Println(i,card)
	}
}

func deal(d deck, handSize int)(deck, deck){
	return d[:handSize], d[handSize:]
}

func (d deck) toString() string {
	// https://pkg.go.dev/strings#example-Join
	return strings.Join([]string(d), ",")
}