package main

func main() {
	// Slices
	cards := deck{newCard(), "This a old card", newCard()}
	cards = append(cards, "This is a latest card")

	// For loops
	cards.print()

}

func newCard() string {
	return "This is a New card"
}
