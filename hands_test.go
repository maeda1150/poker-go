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
)

func TestOnePair(t *testing.T) {
	result := onePair(exOnePair)
	if result == false {
		t.Fatal("this is not one pair.")
	}

	result = onePair(exTwoPair)
	if result == true {
		t.Fatal("this is one pair.")
	}

	result = onePair(exThreeOfAKind)
	if result == true {
		t.Fatal("this is one pair.")
	}

	result = onePair(exStraight)
	if result == true {
		t.Fatal("this is one pair.")
	}

	result = onePair(exFlush)
	if result == true {
		t.Fatal("this is one pair.")
	}
}

func TestTwoPair(t *testing.T) {
	result := twoPair(exOnePair)
	if result == true {
		t.Fatal("this is two pair.")
	}

	result = twoPair(exTwoPair)
	if result == false {
		t.Fatal("this is not two pair.")
	}

	result = twoPair(exThreeOfAKind)
	if result == true {
		t.Fatal("this is two pair.")
	}

	result = twoPair(exStraight)
	if result == true {
		t.Fatal("this is two pair.")
	}

	result = twoPair(exFlush)
	if result == true {
		t.Fatal("this is two pair.")
	}
}

func TestThreeOfAKind(t *testing.T) {
	result := threeOfAKind(exOnePair)
	if result == true {
		t.Fatal("this is three of a kind.")
	}

	result = threeOfAKind(exTwoPair)
	if result == true {
		t.Fatal("this is three of a kind.")
	}

	result = threeOfAKind(exThreeOfAKind)
	if result == false {
		t.Fatal("this is not three of a kind.")
	}

	result = threeOfAKind(exStraight)
	if result == true {
		t.Fatal("this is three of a kind.")
	}

	result = threeOfAKind(exFlush)
	if result == true {
		t.Fatal("this is three of a kind.")
	}
}

func TestStraight(t *testing.T) {
	result := straight(exOnePair)
	if result == true {
		t.Fatal("this is straight.")
	}

	result = straight(exTwoPair)
	if result == true {
		t.Fatal("this is straight.")
	}

	result = straight(exThreeOfAKind)
	if result == true {
		t.Fatal("this is straight.")
	}

	result = straight(exStraight)
	if result == false {
		t.Fatal("this is not straight.")
	}

	result = straight(exRoyalStraight)
	if result == false {
		t.Fatal("this is not straight.")
	}

	result = straight(exFlush)
	if result == true {
		t.Fatal("this is straight.")
	}
}

func TestFlush(t *testing.T) {
	result := flush(exOnePair)
	if result == true {
		t.Fatal("this is flush.")
	}

	result = flush(exTwoPair)
	if result == true {
		t.Fatal("this is flush.")
	}

	result = flush(exThreeOfAKind)
	if result == true {
		t.Fatal("this is flush.")
	}

	result = flush(exStraight)
	if result == true {
		t.Fatal("this is flush.")
	}

	result = flush(exFlush)
	if result == false {
		t.Fatal("this is not flush.")
	}
}
