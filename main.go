package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"testing"
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

	tryTimes := 10000
	if len(os.Args) >= 2 {
		times, err := strconv.Atoi(os.Args[2])
		if err != nil {
			fmt.Println("times is not signed int.")
			os.Exit(1)
		}
		tryTimes = times
	}

	resultCount := NewResultCount()

	bench := testing.Benchmark(func(b *testing.B) {

		results := []Result{}
		for t := 0; t < tryTimes; t++ {
			result := playPreFlop(hands)
			results = append(results, result)

		}

		for _, result := range results {
			if result.IsFourOfAKind {
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
	})

	fmt.Printf("OnePair      : %d ( %f percent )\n", resultCount.CountOnePair, float64(resultCount.CountOnePair)/float64(tryTimes)*100)
	fmt.Printf("TwoPair      : %d ( %f percent )\n", resultCount.CountTwoPair, float64(resultCount.CountTwoPair)/float64(tryTimes)*100)
	fmt.Printf("ThreeOfAKind : %d ( %f percent )\n", resultCount.CountThreeOfAKind, float64(resultCount.CountThreeOfAKind)/float64(tryTimes)*100)
	fmt.Printf("Straight     : %d ( %f percent )\n", resultCount.CountStraight, float64(resultCount.CountStraight)/float64(tryTimes)*100)
	fmt.Printf("Flush        : %d ( %f percent )\n", resultCount.CountFlush, float64(resultCount.CountFlush)/float64(tryTimes)*100)
	fmt.Printf("FullHouse    : %d ( %f percent )\n", resultCount.CountFullHouse, float64(resultCount.CountFullHouse)/float64(tryTimes)*100)
	fmt.Printf("FourOfAKind  : %d ( %f percent )\n", resultCount.CountFourOfAKind, float64(resultCount.CountFourOfAKind)/float64(tryTimes)*100)
	fmt.Println(strconv.FormatFloat(bench.T.Seconds(), 'f', 20, 64) + " sec")
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
