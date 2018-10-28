package main

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
