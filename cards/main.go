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
	hand, remainingDeck := deal(cards2, 3) //5
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
	// Error handling
	// testRf := newDeckFromFile("test non-exist file")
	// testRf.print()
	fmt.Println("======")

	fmt.Println("=== Shuffle Decks ===")
	fmt.Println("= Normal shuffle (get same result every time) =")
	cardsForNormalShuffle := newDeck()
	cardsForNormalShuffle.normalShuffle()
	cardsForNormalShuffle.print()
	fmt.Println("= Truly Random Shuffle =")
	cardsForTrulyRandomShuffle := newDeck()
	cardsForTrulyRandomShuffle.trulyRandomShuffle()
	cardsForTrulyRandomShuffle.print()
	fmt.Println("======")

}

func newCard() string {
	return "This is a New card"
}
