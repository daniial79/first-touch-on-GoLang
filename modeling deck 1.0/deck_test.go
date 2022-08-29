package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	d := newDeck()

	if len(d) != 52 {
		t.Errorf("Expected deck length of 52, But got %v", len(d))
	}

	if d[0] != "Ace of Spades" {
		t.Errorf("Expected Ace of Spades at index zero, But got %v", d[0])
	}

	if d[len(d)-1] != "King of Clubs" {
		t.Errorf("Expected King of Clubs at last index, But got %v", d[len(d)-1])
	}
}

func TestSaveToDeckAndNewDeckFromFile(t *testing.T) {
	os.Remove("_decktesting")
	d := newDeck()
	d.saveToFile("_decktesting")
	loadedDeck := newDeckFromFile("_decktesting")

	if len(loadedDeck) != 52 {
		t.Errorf("Expected 52 cards, But got %v", len(loadedDeck))
	}
	os.Remove("_decktesting")
}

func TestDeal(t *testing.T) {
	d := newDeck()
	hand, remainingDeck := deal(d, 10)

	if hand == nil {
		t.Error("Expected deck type for hand, Got nil")
	}

	if remainingDeck == nil {
		t.Error("Expected deck type for remainingDeck, Got nil")
	}

	if len(hand) != 10 {
		t.Errorf("Expected hand length of 10, But got %v", len(hand))
	}

	if len(remainingDeck) != 42 {
		t.Errorf("Expected remainingDeck length of 10, But got %v", len(remainingDeck))
	}
}
