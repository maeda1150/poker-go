package main

import (
	"fmt"
	"os"
	"sync"
	"time"

	"github.com/notnil/combos"
)

func main() {
	start := time.Now()

	suits, nums := splitSuitsAndNumbers(os.Args[1])

	if (len(suits) != 2) || (len(nums) != 2) {
		fmt.Println("invalid card length")
		os.Exit(1)
	}

	hands := createCardsFromSuitsAndNumbers(suits, nums)
	fmt.Println(hands)

	deck := createDeck()
	deck = removeCardsFromDeck(deck, hands)
	combs := combos.New(len(deck), 5)
	tryTimes := len(combs)
	combsList := splitCombs(combs)
	resultCount := NewResultCount()

	wg := new(sync.WaitGroup)
	results := make(chan []Result, len(combsList))
	for i, coms := range combsList {
		wg.Add(1)
		go func(hands []Card, deck []Card, coms [][]int, i int) {
			defer wg.Done()
			playPreFlopWithTimes(hands, deck, coms, results, i)
		}(hands, deck, coms, i)
	}
	go func() {
		wg.Wait()
		close(results)
	}()

	rs := []Result{}
	for result := range results {
		for _, r := range result {
			rs = append(rs, r)
		}
	}

	resultCount = calcResultCount(rs)

	fmt.Printf("OnePair        : %d ( %f %v ), with hands : %d ( %f %v )\n", resultCount.CountOnePair, float64(resultCount.CountOnePair)/float64(tryTimes)*100, "%", resultCount.CountOnePairWithHands, float64(resultCount.CountOnePairWithHands)/float64(tryTimes)*100, "%")
	fmt.Printf("TwoPair        : %d ( %f %v ), with hands : %d ( %f %v )\n", resultCount.CountTwoPair, float64(resultCount.CountTwoPair)/float64(tryTimes)*100, "%", resultCount.CountTwoPairWithHands, float64(resultCount.CountTwoPairWithHands)/float64(tryTimes)*100, "%")
	fmt.Printf("ThreeOfAKind   : %d ( %f %v ), with hands : %d ( %f %v )\n", resultCount.CountThreeOfAKind, float64(resultCount.CountThreeOfAKind)/float64(tryTimes)*100, "%", resultCount.CountThreeOfAKindWithHands, float64(resultCount.CountThreeOfAKindWithHands)/float64(tryTimes)*100, "%")
	fmt.Printf("Straight       : %d ( %f %v ), with hands : %d ( %f %v )\n", resultCount.CountStraight, float64(resultCount.CountStraight)/float64(tryTimes)*100, "%", resultCount.CountStraightWithHands, float64(resultCount.CountStraightWithHands)/float64(tryTimes)*100, "%")
	fmt.Printf("Flush          : %d ( %f %v ), with hands : %d ( %f %v )\n", resultCount.CountFlush, float64(resultCount.CountFlush)/float64(tryTimes)*100, "%", resultCount.CountFlushWithHands, float64(resultCount.CountFlushWithHands)/float64(tryTimes)*100, "%")
	fmt.Printf("FullHouse      : %d ( %f %v ), with hands : %d ( %f %v )\n", resultCount.CountFullHouse, float64(resultCount.CountFullHouse)/float64(tryTimes)*100, "%", resultCount.CountFullHouseWithHands, float64(resultCount.CountFullHouseWithHands)/float64(tryTimes)*100, "%")
	fmt.Printf("FourOfAKind    : %d ( %f %v ), with hands : %d ( %f %v )\n", resultCount.CountFourOfAKind, float64(resultCount.CountFourOfAKind)/float64(tryTimes)*100, "%", resultCount.CountFourOfAKindWithHands, float64(resultCount.CountFourOfAKindWithHands)/float64(tryTimes)*100, "%")
	fmt.Printf("StraightFlush  : %d ( %f %v ), with hands : %d ( %f %v )\n", resultCount.CountStraightFlush, float64(resultCount.CountStraightFlush)/float64(tryTimes)*100, "%", resultCount.CountStraightFlushWithHands, float64(resultCount.CountStraightFlushWithHands)/float64(tryTimes)*100, "%")

	elapsed := time.Since(start)
	fmt.Println(elapsed)
}
