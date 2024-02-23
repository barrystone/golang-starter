package main

import (
	"testing"
)

func TestNewDeck(t *testing.T) {
	d := newDeck()

	if len(d) != 6 {
		t.Errorf("Expected deck length of 6, but got %v", len(d))
	}

	if d[0] != "Ace of Spades" {
		t.Errorf("Expected first card of Ace of Spades, but got %v", d[0])
	}

	if d[len(d)-1] != "Three of Diamonds" {
		t.Errorf("Expected last card of Three of Diamonds, but got %v", d[len(d)-1])
	}
}

func TestSaveToDeckAndNewDeckFromFile(t *testing.T) {
	removeFile("_deckTesting")

	deck := newDeck()
	deck.saveToFile("_deckTesting")

	loadDeck := newDeckFromFile("_deckTesting")

	if len(loadDeck) != 6 {
		t.Errorf("Expected 6 cards in deck, got %v", len(loadDeck))
	}

	removeFile("_deckTesting")
}
