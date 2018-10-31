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

// boad と hands で onePair になっていることが前提
func onePairWithHands(boad []Card, hands []Card) bool {
	if hands[0].Number == hands[1].Number {
		return true
	}
	for _, b := range boad {
		for _, hand := range hands {
			if b.Number == hand.Number {
				return true
			}
		}
	}
	return false
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

// boad と hands で twoPair になっていることが前提
func twoPairWithHands(boad []Card, hands []Card) bool {
	if hands[0].Number == hands[1].Number {
		return true
	}
	for _, b := range boad {
		for _, hand := range hands {
			if b.Number == hand.Number {
				return true
			}
		}
	}
	return false
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

// boad と hands で threeOfAKind になっていることが前提
func threeOfAKindWithHands(boad []Card, hands []Card) bool {
	all := append(hands, boad...)
	numCount := map[int]int{}
	for _, a := range all {
		numCount[a.Number] = numCount[a.Number] + 1
	}
	threeNum := 0
	for k, v := range numCount {
		if v == 3 {
			threeNum = k
		}
	}
	for _, hand := range hands {
		if hand.Number == threeNum {
			return true
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

// boad と hands で straight になっていることが前提
func straightWithHands(boad []Card, hands []Card) bool {
	all := append(hands, boad...)
	comb := combos.New(len(all), 5)
	straights := [][]Card{}
	for _, com := range comb {
		a, b, c, d, e := all[com[0]], all[com[1]], all[com[2]], all[com[3]], all[com[4]]
		abcde := []Card{a, b, c, d, e}
		if straight(abcde) {
			straights = append(straights, abcde)
		}
	}
	for _, s := range straights {
		for _, c := range s {
			if c.Same(hands[0]) || c.Same(hands[1]) {
				return true
			}
		}
	}
	return false
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

// boad と hands で flush になっていることが前提
func flushWithHands(boad []Card, hands []Card) bool {
	all := append(hands, boad...)
	comb := combos.New(len(all), 5)
	flushes := [][]Card{}
	for _, com := range comb {
		a, b, c, d, e := all[com[0]], all[com[1]], all[com[2]], all[com[3]], all[com[4]]
		abcde := []Card{a, b, c, d, e}
		if flush(abcde) {
			flushes = append(flushes, abcde)
		}
	}
	for _, f := range flushes {
		for _, c := range f {
			if c.Same(hands[0]) || c.Same(hands[1]) {
				return true
			}
		}
	}
	return false
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

// boad と hands で fullHouse になっていることが前提
func fullHouseWithHands(boad []Card, hands []Card) bool {
	all := append(hands, boad...)
	comb := combos.New(len(all), 5)
	fullHouses := [][]Card{}
	for _, com := range comb {
		a, b, c, d, e := all[com[0]], all[com[1]], all[com[2]], all[com[3]], all[com[4]]
		abcde := []Card{a, b, c, d, e}
		if fullHouse(abcde) {
			fullHouses = append(fullHouses, abcde)
		}
	}
	for _, f := range fullHouses {
		for _, c := range f {
			if c.Same(hands[0]) || c.Same(hands[1]) {
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

// boad と hands で fourOfAKind になっていることが前提
func fourOfAKindWithHands(boad []Card, hands []Card) bool {
	all := append(hands, boad...)
	numCount := map[int]int{}
	for _, a := range all {
		numCount[a.Number] = numCount[a.Number] + 1
	}
	fourNum := 0
	for k, v := range numCount {
		if v == 4 {
			fourNum = k
		}
	}
	for _, hand := range hands {
		if hand.Number == fourNum {
			return true
		}
	}
	return false
}

func straightFlush(cards []Card) bool {
	return straightImplementation(cards) && flushImplementation(cards)
}

// boad と hands で straightFlush になっていることが前提
func straightFlushWithHands(boad []Card, hands []Card) bool {
	all := append(hands, boad...)
	comb := combos.New(len(all), 5)
	straightFlushes := [][]Card{}
	for _, com := range comb {
		a, b, c, d, e := all[com[0]], all[com[1]], all[com[2]], all[com[3]], all[com[4]]
		abcde := []Card{a, b, c, d, e}
		if straightFlush(abcde) {
			straightFlushes = append(straightFlushes, abcde)
		}
	}
	for _, s := range straightFlushes {
		for _, c := range s {
			if c.Same(hands[0]) || c.Same(hands[1]) {
				return true
			}
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
