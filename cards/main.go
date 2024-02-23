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
	fmt.Println()

	fmt.Println("=== cards2 (with newDeck())===")
	cards2 := newDeck()
	cards2.print()
	fmt.Println("======")
	fmt.Println()

	fmt.Println("=== deal() ===")
	hand, remainingDeck := deal(cards2, 3) //5
	fmt.Println("= hand =")
	hand.print()
	fmt.Println()
	fmt.Println("= remainingDeck =")
	remainingDeck.print()
	fmt.Println("======")
	fmt.Println()

	fmt.Println("=== toStrring ===")
	fmt.Println(cards2.toString())
	fmt.Println("======")
	fmt.Println()

	fmt.Println("=== Save to file & Read from file ===")
	// Save to file
	CARDS_SAVE_1 := "cards save"
	CARDS_SAVE_2 := "cards2 save"
	cards.saveToFile(CARDS_SAVE_1)
	cards2.saveToFile(CARDS_SAVE_2)
	// Read from file
	cardsRf := newDeckFromFile(CARDS_SAVE_1)
	cards2Rf := newDeckFromFile(CARDS_SAVE_2)
	fmt.Println("= Read", CARDS_SAVE_1, "=")
	cardsRf.print()
	fmt.Println()
	fmt.Println("= Read", CARDS_SAVE_2, "=")
	cards2Rf.print()
	// Error handling, *Comment out to see the error
	// testRf := newDeckFromFile("test non-exist file")
	// testRf.print()
	fmt.Println("======")
	fmt.Println()

	fmt.Println("=== Shuffle Decks ===")
	fmt.Println("= Normal shuffle (get same result every time) =")
	cardsForNormalShuffle := newDeck()
	cardsForNormalShuffle.normalShuffle()
	cardsForNormalShuffle.print()
	fmt.Println()
	fmt.Println("= Truly Random Shuffle =")
	cardsForTrulyRandomShuffle := newDeck()
	cardsForTrulyRandomShuffle.trulyRandomShuffle()
	cardsForTrulyRandomShuffle.print()
	fmt.Println("======")
	fmt.Println()
}

func newCard() string {
	return "This is a New card"
}
