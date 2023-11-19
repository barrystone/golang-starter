package main

import "fmt"

func main() {

	fmt.Println("=== cards ===")
	// Slices
	cards := deck{newCard(), "This a old card", newCard()}
	cards = append(cards, "This is a latest card")
	// For loops
	cards.print()
	fmt.Println("======")
	
	fmt.Println("=== cards2 (with newDeck())===")
	cards2 := newDeck()
	cards2.print()
	fmt.Println("======")


	fmt.Println("=== deal() ===")
	hand, remainingDeck := deal(cards2, 5)
	fmt.Println("= hand =")
	hand.print()
	fmt.Println("= remainingDeck =")
	remainingDeck.print()
	fmt.Println("======")

	fmt.Println("=== toStrring ===")
	fmt.Println(cards2.toString())
}

func newCard() string {
	return "This is a New card"
}
