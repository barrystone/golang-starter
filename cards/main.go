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
	fmt.Println("======")

	fmt.Println("=== Save to file & Read from file ===")
	// Save to file
	cards.saveToFile("cards save")
	cards2.saveToFile("cards2 save")
	// Read from file
	cardsRf := newDeckFromFile("cards save")
	cards2Rf := newDeckFromFile("cards save")
	cardsRf.print()
	cards2Rf.print()

	testRf := newDeckFromFile("test non-exist file")
	testRf.print()

}

func newCard() string {
	return "This is a New card"
}
