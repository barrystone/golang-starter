package main

import "fmt"

func main() {

	fmt.Print("=== cards ===")
	// Slices
	cards := deck{newCard(), "This a old card", newCard()}
	cards = append(cards, "This is a latest card")
	// For loops
	cards.print()
	fmt.Print("======")
	
	cards2 := newDeck()
	fmt.Print("=== cards2 (with newDeck())===")
	cards2.print()
	fmt.Print("======")


	fmt.Print("=== deal() ===")
	hand, remainingDeck := deal(cards2, 5)
	fmt.Print("= hand =")
	hand.print()
	fmt.Print("= remainingDeck =")
	remainingDeck.print()
	fmt.Print("======")


}

func newCard() string {
	return "This is a New card"
}
