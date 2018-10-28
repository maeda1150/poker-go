package main

import (
	"fmt"
	"os"
)

func main() {
	suits, nums := splitSuitsAndNumbers(os.Args[1])

	if (len(suits) != 2) || (len(nums) != 2) {
		fmt.Println("invalid card length")
		os.Exit(1)
	}

	hands := createCardsFromSuitsAndNumbers(suits, nums)
	fmt.Println(hands)

	//cards := []string{"1", "2", "3", "4", "4", "6", "7"}
	//result, _ := onePair(cards)
	//fmt.Println(result)

	//cards = []string{"1", "2", "3", "4", "5", "6", "7"}
	//result, _ = onePair(cards)
	//fmt.Println(result)

	//cards = []string{"11", "22", "33", "22", "55", "66", "77"}
	//result, _ = onePair(cards)
	//fmt.Println(result)

	card := Card{"s", 10}
	fmt.Println(card.Valid())

	card = Card{"s", 14}
	fmt.Println(card.Valid())

	card = Card{"v", 10}
	fmt.Println(card.Valid())

	cards := []Card{Card{"s", 10}, Card{"h", 11}, Card{"s", 9}, Card{"s", 7}, Card{"s", 1}}
	result := onePair(cards)
	fmt.Println(result)

	cards = []Card{Card{"s", 10}, Card{"h", 9}, Card{"s", 9}, Card{"s", 7}, Card{"s", 1}}
	result = onePair(cards)
	fmt.Println(result)
}
