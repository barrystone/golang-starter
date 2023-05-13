package main

import "fmt"

func main() {
	// Slices
	cards := []string{newCard(),"This a old card",newCard()} 
	cards = append(cards, "This is a latest card")

	// For loops
	for i, card := range cards {
		fmt.Println(i ,card)
	}

}

func newCard() string {
	return "This is a New card"
}
