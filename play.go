package main

import (
	"math/rand"
	"time"
)

type Result struct {
	IsStraightFlush bool
	IsFourOfAKind   bool
	IsFullHouse     bool
	IsFlush         bool
	IsStraight      bool
	IsThreeOfAKind  bool
	IsTwoPair       bool
	IsOnePair       bool
}

func NewResult() Result {
	result := Result{}
	result.IsStraightFlush = false
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
	CountStraightFlush int
	CountFourOfAKind   int
	CountFullHouse     int
	CountFlush         int
	CountStraight      int
	CountThreeOfAKind  int
	CountTwoPair       int
	CountOnePair       int
}

func NewResultCount() ResultCount {
	resultCount := ResultCount{}
	resultCount.CountStraightFlush = 0
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

	for com := range combinations(all, 5, 1) {
		if straightFlush(com) {
			result.IsStraightFlush = true
		} else if fourOfAKind(com) {
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

func calcResultCount(results chan Result) ResultCount {
	resultCount := NewResultCount()
	for result := range results {
		if result.IsStraightFlush {
			resultCount.CountStraightFlush++
		} else if result.IsFourOfAKind {
			resultCount.CountFourOfAKind++
		} else if result.IsFullHouse {
			resultCount.CountFullHouse++
		} else if result.IsFlush {
			resultCount.CountFlush++
		} else if result.IsStraight {
			resultCount.CountStraight++
		} else if result.IsThreeOfAKind {
			resultCount.CountThreeOfAKind++
		} else if result.IsTwoPair {
			resultCount.CountTwoPair++
		} else if result.IsOnePair {
			resultCount.CountOnePair++
		}
	}
	return resultCount
}
