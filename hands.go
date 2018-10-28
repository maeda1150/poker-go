package main

import (
	"sort"

	"github.com/notnil/combos"
)

func onePair(cards []Card) bool {
	comb := combos.New(5, 2)
	count := 0
	for _, com := range comb {
		if cards[com[0]].Number == cards[com[1]].Number {
			count++
		}
	}
	return count == 1
}

func twoPair(cards []Card) bool {
	comb := combos.New(5, 2)
	set := [][]Card{}
	for _, com := range comb {
		if cards[com[0]].Number == cards[com[1]].Number {
			set = append(set, []Card{cards[com[0]], cards[com[1]]})
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
	comb := combos.New(5, 3)
	for _, com := range comb {
		a := cards[com[0]]
		b := cards[com[1]]
		c := cards[com[2]]
		if a.Number == b.Number && a.Number == c.Number {
			rest := restOfCards(cards, []Card{a, b, c})
			if rest[0].Number != rest[1].Number {
				return true
			}
		}
	}
	return false
}

func straight(cards []Card) bool {
	if straightFlush(cards) {
		return false
	}
	return straightImplementation(cards)
}

func straightImplementation(cards []Card) bool {
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
	if straightFlush(cards) {
		return false
	}
	return flushImplementation(cards)
}

func flushImplementation(cards []Card) bool {
	return cards[0].Suit == cards[1].Suit && cards[0].Suit == cards[2].Suit && cards[0].Suit == cards[3].Suit && cards[0].Suit == cards[4].Suit
}

func fullHouse(cards []Card) bool {
	comb := combos.New(5, 3)
	for _, com := range comb {
		a := cards[com[0]]
		b := cards[com[1]]
		c := cards[com[2]]
		if a.Number == b.Number && a.Number == c.Number {
			rest := restOfCards(cards, []Card{a, b, c})
			if a.Number != rest[0].Number && rest[0].Number == rest[1].Number {
				return true
			}
		}
	}
	return false
}

func fourOfAKind(cards []Card) bool {
	comb := combos.New(5, 4)
	for _, com := range comb {
		a := cards[com[0]]
		b := cards[com[1]]
		c := cards[com[2]]
		d := cards[com[3]]
		if a.Number == b.Number && a.Number == c.Number && a.Number == d.Number {
			return true
		}
	}
	return false
}

func straightFlush(cards []Card) bool {
	return straightImplementation(cards) && flushImplementation(cards)
}

func onePairWithHands(cards []Card, hands []Card) bool {
	if !onePair(cards) {
		return false
	}
	for _, card := range cards {
		if card.Same(hands[0]) || card.Same(hands[1]) {
			return true
		}
	}
	return false
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
