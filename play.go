package main

import (
	"math/rand"
	"time"
)

type Result struct {
	IsFourOfAKind  bool
	IsFullHouse    bool
	IsFlush        bool
	IsStraight     bool
	IsThreeOfAKind bool
	IsTwoPair      bool
	IsOnePair      bool
}

func NewResult() Result {
	result := Result{}
	result.IsFourOfAKind = false
	result.IsFullHouse = false
	result.IsFlush = false
	result.IsStraight = false
	result.IsThreeOfAKind = false
	result.IsTwoPair = false
	result.IsOnePair = false
	return result
}

type ResultCount struct {
	CountFourOfAKind  int
	CountFullHouse    int
	CountFlush        int
	CountStraight     int
	CountThreeOfAKind int
	CountTwoPair      int
	CountOnePair      int
}

func NewResultCount() ResultCount {
	resultCount := ResultCount{}
	resultCount.CountFourOfAKind = 0
	resultCount.CountFullHouse = 0
	resultCount.CountFlush = 0
	resultCount.CountStraight = 0
	resultCount.CountThreeOfAKind = 0
	resultCount.CountTwoPair = 0
	resultCount.CountOnePair = 0
	return resultCount
}

func playPreFlop(hands []Card) Result {
	result := NewResult()

	deck := createDeck()
	deck = removeCardsFromDeck(deck, hands)
	shuffleDeck(deck)
	boad := []Card{}
	for i := 0; i < 5; i++ {
		rand.Seed(time.Now().UnixNano())
		boad = append(boad, deck[rand.Intn(50)])
	}
	all := append(hands, boad...)

	for com := range combinations(all, 5, 10) {
		if fourOfAKind(com) {
			result.IsFourOfAKind = true
		} else if fullHouse(com) {
			result.IsFullHouse = true
		} else if flush(com) {
			result.IsFlush = true
		} else if straight(com) {
			result.IsStraight = true
		} else if threeOfAKind(com) {
			result.IsThreeOfAKind = true
		} else if twoPair(com) {
			result.IsTwoPair = true
		} else if onePair(com) {
			result.IsOnePair = true
		}
	}
	return result
}
