package main

import (
	"testing"
	"os"
)

// Test Function usually has an uppercase
func TestNewDeck(t *testing.T) {
	d := newDeck()

	if len(d) != 16 {
		t.Errorf("Expected deck length of 16, but got %v", len(d))
	}

	if d[0] != "Spades of Ace" {
		t.Errorf("Expected first card of Ace of Spades, but got %v", d[0])
	}

	if d[len(d) -1] != "Clubs of Four" {
		t.Errorf("Expected last card of Four of Clubs, but got %v", d[len(d) - 1])
	}

}

func TestSaveToDeckAndNewDeckTestFromFile (t *testing.T) {
	os.Remove("_decktesting")

	deck := newDeck()
	deck.saveToFile("_decktesting")

	loadedDeck := newDeckFromFile("_decktesting")

	if len(loadedDeck) != 16 {
		t.Errorf("Expected 16 cards in deck, but got %v", len(loadedDeck))
	}

	os.Remove("_decktesting")
}




/*
To make a test , create a new file ending in _test.go

To run all tests in a package, run the command
- go test
*/