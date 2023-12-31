package main

import (
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
