package main

type Card struct {
	Suit   string
	Number int
}

func (card Card) Same(c Card) bool {
	return card.Number == c.Number && card.Suit == c.Suit
}

func (card Card) Valid() bool {
	return card.validSuit() && card.validNumber()
}

func (card Card) validNumber() bool {
	return contains([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13}, card.Number)
}

func (card Card) validSuit() bool {
	suits := []string{"s", "h", "d", "c"}
	for _, suit := range suits {
		if suit == card.Suit {
			return true
		}
	}
	return false
}

func contains(s []int, e int) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}
