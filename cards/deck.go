package main

import (
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

// Create a new type of 'deck'
// which is a slice of strings
type deck []string

func newDeck() deck {
	cards := deck{}

	cardSuits := []string{"Spades", "Diamods"} //  "Hearts", "Clubs"

	cardValues := []string{"Ace", "Two", "Three"} // "Four"

	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, value+" of "+suit)
		}
	}

	return cards
}

// Receiver function
// 'd' is a variable(like 'this' or 'self'),
// exactly like cards put in there
func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

func (d deck) toString() string {
	return strings.Join([]string(d), ",")
}

func (d deck) saveToFile(fileName string) error {
	return os.WriteFile(fileName, []byte(d.toString()), 0666)
	// "0666" means anyone can read and write this file
}

func newDeckFromFile(fileName string) deck {
	bs, err := os.ReadFile(fileName)
	if err != nil {
		// Option #1 - log the error and return a call to newDeck()
		// Option #2 - log the error and entirely quit the program
		fmt.Println("Error: ", err)
		os.Exit(1)
	}
	//string(bs): Ace of Spades,Two of Spades,Three of Spades, ...
	s := strings.Split(string(bs), ",")
	return deck(s)
}

func (d deck) normalShuffle() {
	for i := range d {
		newPosition := rand.Intn(len(d) - 1)
		// Swap the position of the cards
		d[i], d[newPosition] = d[newPosition], d[i]
	}
}

func (d deck) trulyRandomShuffle() {
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)

	for i := range d {
		newPosition := r.Intn(len(d) - 1)
		d[i], d[newPosition] = d[newPosition], d[i]
	}
}
