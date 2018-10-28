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
