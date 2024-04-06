package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	d := newDeck()

	//4 suits 4 values
	if len(d) != 16 {
		t.Errorf("Expected deck length of 16, but got %v", len(d))
	}
	//First card
	if d[0] != "Ace of Spades" {
		t.Errorf("Expected first card of Ace of Spades, but got %v", d[0])
	}

	if d[len(d)-1] != "Four of Clubs" {
		t.Errorf("Expected last card of Four of Clubs, but got %v", d[len(d)-1])
	}
}

func TestSaveToDeckAndNewDeckTestFromFile(t *testing.T) {
	//Remove any previous deck testing files
	os.Remove("_decktesting")
	//Create a new deck
	deck := newDeck()
	//Save this deck to the file
	deck.saveToFile("_decktesting")
	//Load from disk
	loadedDeck := newDeckFromFile("_decktesting")
	//Write an assertion
	if len(loadedDeck) != 16 {
		t.Errorf("Expected deck length of 16, but got %v", len(loadedDeck))
	}
	//Clean up any files
	os.Remove("_decktesting")
}
