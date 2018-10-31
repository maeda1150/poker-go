package main

import (
	"testing"
)

func TestFindHand(t *testing.T) {
	result := NewResult()
	exOnePairBoad := []Card{{"s", 5}, {"d", 8}, {"s", 10}, {"c", 2}, {"h", 8}, {"d", 9}, {"c", 13}}
	findHand(exOnePairBoad, &result)

	if !result.IsOnePair {
		t.Fatal("expect result is onePair.")
	}
	if result.IsTwoPair {
		t.Fatal("expect result is not TwoPair.")
	}
	if result.IsThreeOfAKind {
		t.Fatal("expect result is not threeOfAKind.")
	}
	if result.IsStraight {
		t.Fatal("expect result is not straight.")
	}
	if result.IsFlush {
		t.Fatal("expect result is not flush.")
	}
	if result.IsFullHouse {
		t.Fatal("expect result is not fullHouse.")
	}
	if result.IsFourOfAKind {
		t.Fatal("expect result is not fourOfAKind.")
	}
	if result.IsStraightFlush {
		t.Fatal("expect result is not straightFlush.")
	}

	result = NewResult()
	exTwoPairBoad := []Card{{"s", 13}, {"d", 8}, {"s", 10}, {"c", 2}, {"h", 8}, {"d", 9}, {"c", 13}}
	findHand(exTwoPairBoad, &result)

	// 組み合わせによっては onePair になるのでヒットする
	if !result.IsOnePair {
		t.Fatal("expect result is onePair.")
	}
	if !result.IsTwoPair {
		t.Fatal("expect result is TwoPair.")
	}
	if result.IsThreeOfAKind {
		t.Fatal("expect result is not threeOfAKind.")
	}
	if result.IsStraight {
		t.Fatal("expect result is not straight.")
	}
	if result.IsFlush {
		t.Fatal("expect result is not flush.")
	}
	if result.IsFullHouse {
		t.Fatal("expect result is not fullHouse.")
	}
	if result.IsFourOfAKind {
		t.Fatal("expect result is not fourOfAKind.")
	}
	if result.IsStraightFlush {
		t.Fatal("expect result is not straightFlush.")
	}

	result = NewResult()
	exThreeOfAKindBoad := []Card{{"s", 13}, {"d", 8}, {"s", 10}, {"c", 2}, {"h", 8}, {"d", 9}, {"c", 8}}
	findHand(exThreeOfAKindBoad, &result)

	// 組み合わせによっては onePair になるのでヒットする
	if !result.IsOnePair {
		t.Fatal("expect result is onePair.")
	}
	if result.IsTwoPair {
		t.Fatal("expect result is TwoPair.")
	}
	if !result.IsThreeOfAKind {
		t.Fatal("expect result is threeOfAKind.")
	}
	if result.IsStraight {
		t.Fatal("expect result is not straight.")
	}
	if result.IsFlush {
		t.Fatal("expect result is not flush.")
	}
	if result.IsFullHouse {
		t.Fatal("expect result is not fullHouse.")
	}
	if result.IsFourOfAKind {
		t.Fatal("expect result is not fourOfAKind.")
	}
	if result.IsStraightFlush {
		t.Fatal("expect result is not straightFlush.")
	}

	result = NewResult()
	exStraightBoad := []Card{{"s", 13}, {"d", 8}, {"s", 10}, {"c", 2}, {"h", 7}, {"d", 9}, {"c", 6}}
	findHand(exStraightBoad, &result)

	if result.IsOnePair {
		t.Fatal("expect result is not onePair.")
	}
	if result.IsTwoPair {
		t.Fatal("expect result is not TwoPair.")
	}
	if result.IsThreeOfAKind {
		t.Fatal("expect result is not threeOfAKind.")
	}
	if !result.IsStraight {
		t.Fatal("expect result is straight.")
	}
	if result.IsFlush {
		t.Fatal("expect result is not flush.")
	}
	if result.IsFullHouse {
		t.Fatal("expect result is not fullHouse.")
	}
	if result.IsFourOfAKind {
		t.Fatal("expect result is not fourOfAKind.")
	}
	if result.IsStraightFlush {
		t.Fatal("expect result is not straightFlush.")
	}

	result = NewResult()
	exFlushBoad := []Card{{"s", 13}, {"d", 8}, {"s", 10}, {"d", 2}, {"d", 3}, {"d", 9}, {"d", 12}}
	findHand(exFlushBoad, &result)

	if result.IsOnePair {
		t.Fatal("expect result is not onePair.")
	}
	if result.IsTwoPair {
		t.Fatal("expect result is not TwoPair.")
	}
	if result.IsThreeOfAKind {
		t.Fatal("expect result is not threeOfAKind.")
	}
	if result.IsStraight {
		t.Fatal("expect result is not straight.")
	}
	if !result.IsFlush {
		t.Fatal("expect result is flush.")
	}
	if result.IsFullHouse {
		t.Fatal("expect result is not fullHouse.")
	}
	if result.IsFourOfAKind {
		t.Fatal("expect result is not fourOfAKind.")
	}
	if result.IsStraightFlush {
		t.Fatal("expect result is not straightFlush.")
	}

	result = NewResult()
	exFullHouseBoad := []Card{{"s", 13}, {"d", 8}, {"s", 10}, {"c", 8}, {"s", 8}, {"d", 9}, {"d", 13}}
	findHand(exFullHouseBoad, &result)

	// 組み合わせによっては onePair になるのでヒットする
	if !result.IsOnePair {
		t.Fatal("expect result is onePair.")
	}
	// 組み合わせによっては twoPair になるのでヒットする
	if !result.IsTwoPair {
		t.Fatal("expect result is TwoPair.")
	}
	// 組み合わせによっては threeOfAKind になるのでヒットする
	if !result.IsThreeOfAKind {
		t.Fatal("expect result is threeOfAKind.")
	}
	if result.IsStraight {
		t.Fatal("expect result is not straight.")
	}
	if result.IsFlush {
		t.Fatal("expect result is not flush.")
	}
	if !result.IsFullHouse {
		t.Fatal("expect result is fullHouse.")
	}
	if result.IsFourOfAKind {
		t.Fatal("expect result is not fourOfAKind.")
	}
	if result.IsStraightFlush {
		t.Fatal("expect result is not straightFlush.")
	}

	result = NewResult()
	exFourOfAKindBoad := []Card{{"s", 1}, {"d", 8}, {"s", 10}, {"c", 8}, {"s", 8}, {"d", 9}, {"h", 8}}
	findHand(exFourOfAKindBoad, &result)

	// 組み合わせによっては onePair になるのでヒットする
	if !result.IsOnePair {
		t.Fatal("expect result is onePair.")
	}
	if result.IsTwoPair {
		t.Fatal("expect result is not TwoPair.")
	}
	// 組み合わせによっては threeOfAKind になるのでヒットする
	if !result.IsThreeOfAKind {
		t.Fatal("expect result is threeOfAKind.")
	}
	if result.IsStraight {
		t.Fatal("expect result is not straight.")
	}
	if result.IsFlush {
		t.Fatal("expect result is not flush.")
	}
	if result.IsFullHouse {
		t.Fatal("expect result is not fullHouse.")
	}
	if !result.IsFourOfAKind {
		t.Fatal("expect result is fourOfAKind.")
	}
	if result.IsStraightFlush {
		t.Fatal("expect result is not straightFlush.")
	}

	result = NewResult()
	exStraightFlushBoad := []Card{{"d", 12}, {"d", 8}, {"d", 10}, {"c", 3}, {"s", 1}, {"d", 9}, {"d", 11}}
	findHand(exStraightFlushBoad, &result)

	if result.IsOnePair {
		t.Fatal("expect result is not onePair.")
	}
	if result.IsTwoPair {
		t.Fatal("expect result is not TwoPair.")
	}
	if result.IsThreeOfAKind {
		t.Fatal("expect result is not threeOfAKind.")
	}
	if result.IsStraight {
		t.Fatal("expect result is not straight.")
	}
	if result.IsFlush {
		t.Fatal("expect result is not flush.")
	}
	if result.IsFullHouse {
		t.Fatal("expect result is not fullHouse.")
	}
	if result.IsFourOfAKind {
		t.Fatal("expect result is not fourOfAKind.")
	}
	if !result.IsStraightFlush {
		t.Fatal("expect result is straightFlush.")
	}
}
