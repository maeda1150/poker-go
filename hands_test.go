package main

import (
	"fmt"
	"testing"
)

func TestOnePair(t *testing.T) {
	cards := []Card{Card{"s", 1}, Card{"s", 2}, Card{"s", 3}, Card{"s", 4}, Card{"h", 4}}
	result, _ := onePair(cards)
	if result == false {
		fmt.Println(cards)
		t.Fatal("expect one pair, this is not one pair.")
	}

	cards = []Card{Card{"s", 1}, Card{"s", 2}, Card{"s", 3}, Card{"s", 4}, Card{"h", 5}}
	result, _ = onePair(cards)
	if result == true {
		fmt.Println(cards)
		t.Fatal("expect not one pair, this is one pair.")
	}

	cards = []Card{Card{"s", 1}, Card{"h", 1}, Card{"s", 3}, Card{"d", 3}, Card{"h", 5}}
	result, _ = onePair(cards)
	if result == true {
		fmt.Println(cards)
		t.Fatal("expect not one pair, this is one pair.")
	}
}
