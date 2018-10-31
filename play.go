package main

import (
	"github.com/notnil/combos"
)

type Result struct {
	IsStraightFlush         bool
	IsFourOfAKind           bool
	IsFullHouse             bool
	IsFlush                 bool
	IsStraight              bool
	IsThreeOfAKind          bool
	IsTwoPair               bool
	IsOnePair               bool
	IsOnePairWithHands      bool
	IsTwoPairWithHands      bool
	IsThreeOfAKindWithHands bool
	IsStraightWithHands     bool
	IsFlushWithHands        bool
	IsFullHouseWithHands    bool
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
	result.IsOnePairWithHands = false
	result.IsTwoPairWithHands = false
	result.IsThreeOfAKindWithHands = false
	result.IsStraightWithHands = false
	result.IsFlushWithHands = false
	result.IsFullHouseWithHands = false
	return result
}

type ResultCount struct {
	CountStraightFlush         int
	CountFourOfAKind           int
	CountFullHouse             int
	CountFlush                 int
	CountStraight              int
	CountThreeOfAKind          int
	CountTwoPair               int
	CountOnePair               int
	CountOnePairWithHands      int
	CountTwoPairWithHands      int
	CountThreeOfAKindWithHands int
	CountStraightWithHands     int
	CountFlushWithHands        int
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
	resultCount.CountOnePairWithHands = 0
	resultCount.CountTwoPairWithHands = 0
	resultCount.CountThreeOfAKindWithHands = 0
	resultCount.CountStraightWithHands = 0
	resultCount.CountFlushWithHands = 0
	return resultCount
}

func playPreFlop(hands []Card) Result {
	result := NewResult()

	deck := createDeck()
	deck = removeCardsFromDeck(deck, hands)
	shuffleDeck(deck)
	boad := []Card{deck[1], deck[3], deck[4], deck[5], deck[7], deck[9]}
	all := append(hands, boad...)

	// 1 つの result に複数の結果が入るようになっているが、
	// calcResultCount で集計時は高い手役のみカウントするので
	// 結果の正当性を保っている。
	comb := combos.New(len(all), 5)
	for _, com := range comb {
		a, b, c, d, e := all[com[0]], all[com[1]], all[com[2]], all[com[3]], all[com[4]]
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

	if result.IsFullHouse {
		result.IsFullHouseWithHands = fullHouseWithHands(boad, hands)
	} else if result.IsFlush {
		result.IsFlushWithHands = flushWithHands(boad, hands)
	} else if result.IsStraight {
		result.IsStraightWithHands = straightWithHands(boad, hands)
	} else if result.IsThreeOfAKind {
		result.IsThreeOfAKindWithHands = threeOfAKindWithHands(boad, hands)
	} else if result.IsTwoPair {
		result.IsTwoPairWithHands = twoPairWithHands(boad, hands)
	} else if result.IsOnePair {
		result.IsOnePairWithHands = onePairWithHands(boad, hands)
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
			if result.IsFlushWithHands {
				resultCount.CountFlushWithHands++
			}
		} else if result.IsStraight {
			resultCount.CountStraight++
			if result.IsStraightWithHands {
				resultCount.CountStraightWithHands++
			}
		} else if result.IsThreeOfAKind {
			resultCount.CountThreeOfAKind++
			if result.IsThreeOfAKindWithHands {
				resultCount.CountThreeOfAKindWithHands++
			}
		} else if result.IsTwoPair {
			resultCount.CountTwoPair++
			if result.IsTwoPairWithHands {
				resultCount.CountTwoPairWithHands++
			}
		} else if result.IsOnePair {
			resultCount.CountOnePair++
			if result.IsOnePairWithHands {
				resultCount.CountOnePairWithHands++
			}
		}
	}
	return resultCount
}
