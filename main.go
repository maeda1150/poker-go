package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"
)

func main() {
	suits, nums := splitSuitsAndNumbers(os.Args[1])

	if (len(suits) != 2) || (len(nums) != 2) {
		fmt.Println("invalid card length")
		os.Exit(1)
	}

	hands := createCardsFromSuitsAndNumbers(suits, nums)
	fmt.Println(hands)

	tryTimes := 1000

	countFourOfAKind := 0
	countFullHouse := 0
	countFlush := 0
	countStraight := 0
	countThreeOfAKind := 0
	countTwoPair := 0
	countOnePair := 0

	for t := 0; t < tryTimes; t++ {
		deck := createDeck()
		deck = removeCardsFromDeck(deck, hands)
		shuffleDeck(deck)
		boad := []Card{}
		for i := 0; i < 5; i++ {
			rand.Seed(time.Now().UnixNano())
			boad = append(boad, deck[rand.Intn(50)])
		}
		all := append(hands, boad...)

		isFourOfAKind := false
		isFullHouse := false
		isFlush := false
		isStraight := false
		isThreeOfAKind := false
		isTwoPair := false
		isOnePair := false

		for com := range combinations(all, 5, 20) {
			if fourOfAKind(com) {
				isFourOfAKind = true
			} else if fullHouse(com) {
				isFullHouse = true
			} else if flush(com) {
				isFlush = true
			} else if straight(com) {
				isStraight = true
			} else if threeOfAKind(com) {
				isThreeOfAKind = true
			} else if twoPair(com) {
				isTwoPair = true
			} else if onePair(com) {
				isOnePair = true
			}
		}

		if isFourOfAKind {
			countFourOfAKind++
		} else if isFullHouse {
			countFullHouse++
		} else if isFlush {
			countFlush++
		} else if isStraight {
			countStraight++
		} else if isThreeOfAKind {
			countThreeOfAKind++
		} else if isTwoPair {
			countTwoPair++
		} else if isOnePair {
			countOnePair++
		}
	}

	fmt.Printf("OnePair      : %d ( %f percent )\n", countOnePair, float64(countOnePair)/float64(tryTimes)*100)
	fmt.Printf("TwoPair      : %d ( %f percent )\n", countTwoPair, float64(countTwoPair)/float64(tryTimes)*100)
	fmt.Printf("ThreeOfAKind : %d ( %f percent )\n", countThreeOfAKind, float64(countThreeOfAKind)/float64(tryTimes)*100)
	fmt.Printf("Straight     : %d ( %f percent )\n", countStraight, float64(countStraight)/float64(tryTimes)*100)
	fmt.Printf("Flush        : %d ( %f percent )\n", countFlush, float64(countFlush)/float64(tryTimes)*100)
	fmt.Printf("FullHouse    : %d ( %f percent )\n", countFullHouse, float64(countFullHouse)/float64(tryTimes)*100)
	fmt.Printf("FourOfAKind  : %d ( %f percent )\n", countFourOfAKind, float64(countFourOfAKind)/float64(tryTimes)*100)
}
