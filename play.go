package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/notnil/combos"
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

func playPreFlopWithTimes(hands []Card, times int, results chan<- []Result, i int) {
	fmt.Printf("start playPreFlopWithTimes --- time: %d, i: %d \n", times, i)
	start := time.Now()
	rs := []Result{}
	for t := 0; t < times; t++ {
		result := playPreFlop(hands)
		rs = append(rs, result)
	}
	elapsed := time.Since(start)
	fmt.Printf("finish playPreFlopWithTimes --- time: %d, i: %d, cost: %s \n", times, i, elapsed)
	results <- rs
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

	comb := combos.New(len(all), 5)
	for _, com := range comb {
		a := all[com[0]]
		b := all[com[1]]
		c := all[com[2]]
		d := all[com[3]]
		e := all[com[4]]
		abcde := []Card{a, b, c, d, e}
		if straightFlush(abcde) {
			result.IsStraightFlush = true
		} else if fourOfAKind(abcde) {
			result.IsFourOfAKind = true
		} else if fullHouse(abcde) {
			result.IsFullHouse = true
		} else if flush(abcde) {
			result.IsFlush = true
		} else if straight(abcde) {
			result.IsStraight = true
		} else if threeOfAKind(abcde) {
			result.IsThreeOfAKind = true
		} else if twoPair(abcde) {
			result.IsTwoPair = true
		} else if onePair(abcde) {
			result.IsOnePair = true
		}
	}
	return result
}

func calcResultCount(results []Result) ResultCount {
	resultCount := NewResultCount()
	for _, result := range results {
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
