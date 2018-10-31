package main

import (
	"testing"
)

var (
	exOnePair       []Card = []Card{Card{"s", 1}, Card{"s", 2}, Card{"s", 3}, Card{"s", 4}, Card{"h", 4}}
	exTwoPair       []Card = []Card{Card{"s", 1}, Card{"s", 2}, Card{"s", 3}, Card{"h", 1}, Card{"h", 2}}
	exThreeOfAKind  []Card = []Card{Card{"s", 1}, Card{"s", 2}, Card{"s", 3}, Card{"h", 2}, Card{"d", 2}}
	exStraight      []Card = []Card{Card{"s", 7}, Card{"s", 6}, Card{"s", 4}, Card{"h", 8}, Card{"d", 5}}
	exRoyalStraight []Card = []Card{Card{"s", 10}, Card{"c", 13}, Card{"s", 1}, Card{"h", 11}, Card{"d", 12}}
	exFlush         []Card = []Card{Card{"c", 5}, Card{"c", 13}, Card{"c", 2}, Card{"c", 7}, Card{"c", 9}}
	exFullHouse     []Card = []Card{Card{"c", 5}, Card{"s", 11}, Card{"d", 5}, Card{"h", 5}, Card{"c", 11}}
	exFourOfAKind   []Card = []Card{Card{"c", 5}, Card{"s", 5}, Card{"d", 8}, Card{"h", 5}, Card{"d", 5}}
	exStraightFlush []Card = []Card{Card{"d", 6}, Card{"d", 9}, Card{"d", 7}, Card{"d", 5}, Card{"d", 8}}
)

func TestOnePair(t *testing.T) {
	if !onePair(exOnePair) {
		t.Fatal("expected one pair.")
	}

	if onePair(exTwoPair) {
		t.Fatal("expected not one pair.")
	}

	if onePair(exThreeOfAKind) {
		t.Fatal("expected not one pair.")
	}

	if onePair(exStraight) {
		t.Fatal("expected not one pair.")
	}

	if onePair(exFlush) {
		t.Fatal("expected not one pair.")
	}

	if onePair(exFullHouse) {
		t.Fatal("expected not one pair.")
	}

	if onePair(exFourOfAKind) {
		t.Fatal("expected not one pair.")
	}

	if onePair(exStraightFlush) {
		t.Fatal("expected not one pair.")
	}
}

func TestOnePairWithHands(t *testing.T) {
	// boad と hands で onePair になっていることが前提
	boad := []Card{Card{"s", 1}, Card{"s", 2}, Card{"s", 5}, Card{"s", 4}, Card{"h", 7}}
	hands := []Card{Card{"c", 10}, Card{"d", 5}}
	if !onePairWithHands(boad, hands) {
		t.Fatal("expected onePairWithHands.")
	}

	boad = []Card{Card{"s", 1}, Card{"s", 2}, Card{"s", 5}, Card{"s", 4}, Card{"h", 7}}
	hands = []Card{Card{"c", 10}, Card{"d", 10}}
	if !onePairWithHands(boad, hands) {
		t.Fatal("expected onePairWithHands.")
	}

	boad = []Card{Card{"s", 1}, Card{"s", 2}, Card{"s", 5}, Card{"s", 4}, Card{"h", 2}}
	hands = []Card{Card{"c", 13}, Card{"d", 12}}
	if onePairWithHands(boad, hands) {
		t.Fatal("expected not onePairWithHands.")
	}
}

func TestTwoPair(t *testing.T) {
	if twoPair(exOnePair) {
		t.Fatal("ecpected not two pair.")
	}

	if !twoPair(exTwoPair) {
		t.Fatal("ecpected two pair.")
	}

	if twoPair(exThreeOfAKind) {
		t.Fatal("ecpected not two pair.")
	}

	if twoPair(exStraight) {
		t.Fatal("ecpected not two pair.")
	}

	if twoPair(exFlush) {
		t.Fatal("ecpected not two pair.")
	}

	if twoPair(exFullHouse) {
		t.Fatal("ecpected not two pair.")
	}

	if twoPair(exFourOfAKind) {
		t.Fatal("ecpected not two pair.")
	}

	if twoPair(exStraightFlush) {
		t.Fatal("ecpected not two pair.")
	}
}

func TestTwoPairWithHands(t *testing.T) {
	// boad と hands で twoPair になっていることが前提
	boad := []Card{Card{"s", 1}, Card{"s", 2}, Card{"s", 5}, Card{"s", 4}, Card{"h", 7}}
	hands := []Card{Card{"c", 4}, Card{"d", 5}}
	if !twoPairWithHands(boad, hands) {
		t.Fatal("expected twoPairWithHands.")
	}

	boad = []Card{Card{"s", 1}, Card{"s", 2}, Card{"s", 5}, Card{"s", 7}, Card{"h", 7}}
	hands = []Card{Card{"c", 10}, Card{"d", 2}}
	if !twoPairWithHands(boad, hands) {
		t.Fatal("expected twoPairWithHands.")
	}

	boad = []Card{Card{"s", 1}, Card{"s", 2}, Card{"s", 5}, Card{"s", 4}, Card{"h", 2}}
	hands = []Card{Card{"c", 12}, Card{"d", 12}}
	if !twoPairWithHands(boad, hands) {
		t.Fatal("expected twoPairWithHands.")
	}

	boad = []Card{Card{"s", 1}, Card{"s", 2}, Card{"s", 5}, Card{"d", 1}, Card{"h", 2}}
	hands = []Card{Card{"c", 12}, Card{"d", 13}}
	if twoPairWithHands(boad, hands) {
		t.Fatal("expected not twoPairWithHands.")
	}
}

func TestThreeOfAKind(t *testing.T) {
	if threeOfAKind(exOnePair) {
		t.Fatal("expected not three of a kind.")
	}

	if threeOfAKind(exTwoPair) {
		t.Fatal("expected not three of a kind.")
	}

	if !threeOfAKind(exThreeOfAKind) {
		t.Fatal("expected three of a kind.")
	}

	if threeOfAKind(exStraight) {
		t.Fatal("this is three of a kind.")
	}

	if threeOfAKind(exFlush) {
		t.Fatal("expected not three of a kind.")
	}

	if threeOfAKind(exFullHouse) {
		t.Fatal("expected not three of a kind.")
	}

	if threeOfAKind(exFourOfAKind) {
		t.Fatal("expected not three of a kind.")
	}

	if threeOfAKind(exStraightFlush) {
		t.Fatal("expected not three of a kind.")
	}
}

func TestThreeOfAKindWithHands(t *testing.T) {
	// boad と hands で twoPair になっていることが前提
	boad := []Card{Card{"s", 1}, Card{"s", 2}, Card{"s", 5}, Card{"s", 4}, Card{"h", 7}}
	hands := []Card{Card{"c", 5}, Card{"d", 5}}
	if !threeOfAKindWithHands(boad, hands) {
		t.Fatal("expected threeOfAKindWithHands.")
	}

	boad = []Card{Card{"s", 1}, Card{"s", 2}, Card{"s", 5}, Card{"s", 7}, Card{"h", 2}}
	hands = []Card{Card{"c", 10}, Card{"d", 2}}
	if !threeOfAKindWithHands(boad, hands) {
		t.Fatal("expected threeOfAKindWithHands.")
	}

	boad = []Card{Card{"s", 1}, Card{"s", 5}, Card{"d", 5}, Card{"s", 4}, Card{"h", 5}}
	hands = []Card{Card{"c", 12}, Card{"d", 12}}
	if threeOfAKindWithHands(boad, hands) {
		t.Fatal("expected not threeOfAKindWithHands.")
	}
}

func TestStraight(t *testing.T) {
	if straight(exOnePair) {
		t.Fatal("expected not straight.")
	}

	if straight(exTwoPair) {
		t.Fatal("expected not straight.")
	}

	if straight(exThreeOfAKind) {
		t.Fatal("expected not straight.")
	}

	if !straight(exStraight) {
		t.Fatal("expected straight.")
	}

	if !straight(exRoyalStraight) {
		t.Fatal("expected not straight.")
	}

	if straight(exFlush) {
		t.Fatal("expected not straight.")
	}

	if straight(exFullHouse) {
		t.Fatal("expected not straight.")
	}

	if straight(exFourOfAKind) {
		t.Fatal("expected not straight.")
	}

	if straight(exStraightFlush) {
		t.Fatal("expected not straight.")
	}
}

func TestStraightWithHands(t *testing.T) {
	// boad と hands で straight になっていることが前提
	boad := []Card{Card{"s", 6}, Card{"s", 9}, Card{"s", 5}, Card{"s", 11}, Card{"h", 8}}
	hands := []Card{Card{"c", 5}, Card{"d", 7}}
	if !straightWithHands(boad, hands) {
		t.Fatal("expected straightWithHands.")
	}

	boad = []Card{Card{"s", 6}, Card{"s", 13}, Card{"s", 5}, Card{"s", 11}, Card{"h", 8}}
	hands = []Card{Card{"c", 9}, Card{"d", 7}}
	if !straightWithHands(boad, hands) {
		t.Fatal("expected straightWithHands.")
	}

	boad = []Card{Card{"s", 6}, Card{"s", 9}, Card{"s", 5}, Card{"s", 7}, Card{"h", 8}}
	hands = []Card{Card{"c", 12}, Card{"d", 12}}
	if straightWithHands(boad, hands) {
		t.Fatal("expected not straightWithHands.")
	}
}

func TestFlush(t *testing.T) {
	if flush(exOnePair) {
		t.Fatal("expected not flush.")
	}

	if flush(exTwoPair) {
		t.Fatal("expected not flush.")
	}

	if flush(exThreeOfAKind) {
		t.Fatal("expected not flush.")
	}

	if flush(exStraight) {
		t.Fatal("expected not flush.")
	}

	if !flush(exFlush) {
		t.Fatal("expected flush.")
	}

	if flush(exFullHouse) {
		t.Fatal("expected not flush.")
	}

	if flush(exFourOfAKind) {
		t.Fatal("expected not flush.")
	}

	if flush(exStraightFlush) {
		t.Fatal("expected not flush.")
	}
}

func TestFullHouse(t *testing.T) {
	if fullHouse(exOnePair) {
		t.Fatal("expected not full house.")
	}

	if fullHouse(exTwoPair) {
		t.Fatal("expected not full house.")
	}

	if fullHouse(exThreeOfAKind) {
		t.Fatal("expected not full house.")
	}

	if fullHouse(exStraight) {
		t.Fatal("expected not full house.")
	}

	if fullHouse(exFlush) {
		t.Fatal("expected not full house.")
	}

	if !fullHouse(exFullHouse) {
		t.Fatal("expected full house.")
	}

	if fullHouse(exFourOfAKind) {
		t.Fatal("expected full house.")
	}

	if fullHouse(exStraightFlush) {
		t.Fatal("expected full house.")
	}
}

func TestFourOfAKind(t *testing.T) {
	if fourOfAKind(exOnePair) {
		t.Fatal("expected not four of a kind.")
	}

	if fourOfAKind(exTwoPair) {
		t.Fatal("expected not four of a kind.")
	}

	if fourOfAKind(exThreeOfAKind) {
		t.Fatal("expected not four of a kind.")
	}

	if fourOfAKind(exStraight) {
		t.Fatal("expected not four of a kind.")
	}

	if fourOfAKind(exFlush) {
		t.Fatal("expected not four of a kind.")
	}

	if fourOfAKind(exFullHouse) {
		t.Fatal("expected not four of a kind.")
	}

	if !fourOfAKind(exFourOfAKind) {
		t.Fatal("expected four of a kind.")
	}

	if fourOfAKind(exStraightFlush) {
		t.Fatal("expected four of a kind.")
	}
}

func TestStraightFlush(t *testing.T) {
	if straightFlush(exOnePair) {
		t.Fatal("expected not straight flush.")
	}

	if straightFlush(exTwoPair) {
		t.Fatal("expected not straight flush.")
	}

	if straightFlush(exThreeOfAKind) {
		t.Fatal("expected not straight flush.")
	}

	if straightFlush(exStraight) {
		t.Fatal("expected not straight flush.")
	}

	if straightFlush(exFlush) {
		t.Fatal("expected not straight flush.")
	}

	if straightFlush(exFullHouse) {
		t.Fatal("expected not straight flush.")
	}

	if straightFlush(exFourOfAKind) {
		t.Fatal("expected not straight flush.")
	}

	if !straightFlush(exStraightFlush) {
		t.Fatal("expected straight flush.")
	}
}
