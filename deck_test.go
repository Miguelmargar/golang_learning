package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {

	d := newDeck()

	if len(d) != 36 {
		t.Errorf("Expected deck lenght of 36 but got: %v", len(d))
	}

	if d[0] != "Ace of Spades" {
		t.Errorf("Expected card: Ace of Spades but got: %s", d[0])
	}

	if d[len(d)-1] != "Nine of Clubs" {
		t.Errorf("Expected card: Nine of Clubs but got: %s", d[len(d)-1])
	}
}

func TestSaveToDeckAndNewDeckFromFile(t *testing.T) {

	os.Remove("_decktesting")

	deck := newDeck()
	deck.saveToFile("_decktesting")

	loadedDeck := newDeckFromFile("_decktesting")

	if len(loadedDeck) != 36 {
		t.Errorf("Expected 36 cards in deck, got %v", len(loadedDeck))
	}

	os.Remove("_decktesting")
}

func TestDeal(t *testing.T) {

	d := newDeck()

	d1, d2 := deal(d, 4)

	if len(d1) != 4 {
		t.Errorf("Expected deck to be of lenght 4 and got lenght: %v", len(d1))
	}

	if len(d2) != 32 {
		t.Errorf("Expected deck to be of lenght 32 and got lenght: %v", len(d2))
	}
}
