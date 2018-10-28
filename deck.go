package main

import (
	"math/rand"
	"time"
)

func createDeck() []Card {
	cards := []Card{}
	for _, suit := range []string{"s", "h", "d", "c"} {
		for _, num := range []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13} {
			card := Card{suit, num}
			cards = append(cards, card)
		}
	}
	return cards
}

func removeCardsFromDeck(deck []Card, cards []Card) []Card {
	newDeck := []Card{}
	for _, d := range deck {
		same := false
		for _, card := range cards {
			if d.Same(card) {
				same = true
			}
		}
		if !same {
			newDeck = append(newDeck, d)
		}
	}
	return newDeck
}

func shuffleDeck(deck []Card) {
	rand.Seed(time.Now().UnixNano())
	times := 5 + rand.Intn(5)
	for t := 0; t < times; t++ {
		n := len(deck)
		for i := n - 1; i >= 0; i-- {
			j := rand.Intn(i + 1)
			deck[i], deck[j] = deck[j], deck[i]
		}
	}
}
