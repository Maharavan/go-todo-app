package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	d := newDeck()
	if len(d) != 16 {
		t.Errorf("Length varies %v", len(d))
	}

	if d[0] != "Ace of Spades" {
		t.Errorf("First card varies %s", d[0])
	}

	if d[len(d)-1] != "Four of clubs" {
		t.Errorf("Last card varies %s", d[len(d)-1])
	}
}

func TestSaveToDeckandNewDeckFromFile(t *testing.T) {
	os.Remove("testDeck.txt")
	d := newDeck()
	d.saveTofile("testDeck.txt")
	getFromFile := newDeckFromFile("testDeck.txt")

	if len(getFromFile) != 16 {
		t.Errorf("Expected 16 ,loaded %v", len(getFromFile))
	}
	os.Remove("testDeck.txt")

}
