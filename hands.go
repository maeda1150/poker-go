package main

import (
	"sort"
)

func onePair(cards []Card) bool {
	comb := combinations(cards, 2, 10)
	count := 0
	for com := range comb {
		if com[0].Number == com[1].Number {
			count++
		}
	}
	return count == 1
}

func twoPair(cards []Card) bool {
	comb := combinations(cards, 2, 10)
	set := [][]Card{}
	for com := range comb {
		if com[0].Number == com[1].Number {
			set = append(set, []Card{com[0], com[1]})
		}
	}
	if len(set) != 2 {
		return false
	}
	return duplicatePair(set[0], set[1])
}

func threeOfAKind(cards []Card) bool {
	if fourOfAKind(cards) {
		return false
	}
	comb := combinations(cards, 3, 10)
	for com := range comb {
		if com[0].Number == com[1].Number && com[0].Number == com[2].Number {
			rest := restOfCards(cards, com)
			if rest[0].Number != rest[1].Number {
				return true
			}
		}
	}
	return false
}

func straight(cards []Card) bool {
	sort.Slice(cards, func(i, j int) bool { return cards[i].Number < cards[j].Number })
	royals := []int{1, 10, 11, 12, 13}
	isRoyal := true
	for i, card := range cards {
		if card.Number != royals[i] {
			isRoyal = false
		}
	}
	if isRoyal {
		return true
	}
	baseNum := cards[0].Number
	for i, card := range cards {
		if card.Number != baseNum+i {
			return false
		}
	}
	return true
}

func flush(cards []Card) bool {
	return cards[0].Suit == cards[1].Suit && cards[0].Suit == cards[2].Suit && cards[0].Suit == cards[3].Suit && cards[0].Suit == cards[4].Suit
}

func fullHouse(cards []Card) bool {
	comb := combinations(cards, 3, 10)
	for com := range comb {
		if com[0].Number == com[1].Number && com[0].Number == com[2].Number {
			rest := restOfCards(cards, com)
			if com[0].Number != rest[0].Number && rest[0].Number == rest[1].Number {
				return true
			}
		}
	}
	return false
}

func fourOfAKind(cards []Card) bool {
	comb := combinations(cards, 4, 10)
	for com := range comb {
		if com[0].Number == com[1].Number && com[0].Number == com[2].Number && com[0].Number == com[3].Number {
			return true
		}
	}
	return false
}

func combinations(list []Card, select_num, buf int) (c chan []Card) {
	c = make(chan []Card, buf)
	go func() {
		defer close(c)
		switch {
		case select_num == 0:
			c <- []Card{}
		case select_num == len(list):
			c <- list
		case len(list) < select_num:
			return
		default:
			for i := 0; i < len(list); i++ {
				for sub_comb := range combinations(list[i+1:], select_num-1, buf) {
					c <- append([]Card{list[i]}, sub_comb...)
				}
			}
		}
	}()
	return
}

func duplicatePair(a []Card, b []Card) bool {
	if a[0].Same(b[0]) {
		return false
	}
	if a[0].Same(b[1]) {
		return false
	}
	if a[1].Same(b[0]) {
		return false
	}
	if a[1].Same(b[1]) {
		return false
	}
	return true
}

func restOfCards(a []Card, b []Card) []Card {
	rest := []Card{}
	for _, c := range a {
		duplicate := false
		for _, d := range b {
			if d.Same(c) {
				duplicate = true
			}
		}
		if duplicate {
			continue
		}
		rest = append(rest, c)
	}
	return rest
}
