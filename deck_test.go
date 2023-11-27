package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	d := newDeck()

	if len(d) != 16 {
		t.Errorf("Expected deck length of 16, but got %v", len(d))
	}

	if d[0] != "Spades of Acc" || d[len(d)-1] != "Clubs of Four" {
		t.Errorf("Expected deck first card to be Spades of Acc, and last card to be Clubs of Four. but we get deck first card of %v and last card of %v.", d[0], d[len(d)-1])
	}
}

func TestSaveToDisk(t *testing.T) {
	os.Remove("_decktesting")

	d := newDeck()

	d.saveToFile("_decktesting")

	loadedDeck := readAFile("_decktesting")

	if len(loadedDeck) != 16 {
		t.Errorf("Expected deck length of 16, but got %v", len(d))
	}

	os.Remove("_decktesting")
}
