package main

import (
	"testing"
)

func TestRemoveCardsFromDeck(t *testing.T) {
	deck := createDeck()
	cards := []Card{{"s", 5}, {"d", 8}}
	rest := removeCardsFromDeck(deck, cards)
	for _, r := range rest {
		if r.Same(cards[0]) || r.Same(cards[1]) {
			t.Fatal("Did not remove cards from deck.")
		}
	}
	if len(rest) != 50 {
		t.Fatal("Rest was not 50 cards.")
	}
}

func TestShuffleDeck(t *testing.T) {
	deck1 := createDeck()
	deck2 := createDeck()

	shuffleDeck(deck1)

	if len(deck1) != 52 {
		t.Fatal("Rest was not 52 cards.")
	}

	for _, a := range deck2 {
		match := false
		for _, b := range deck1 {
			if a.Same(b) {
				match = true
			}
		}
		if !match {
			t.Fatal("Shuffled cards something wrong.")
		}
	}
}
