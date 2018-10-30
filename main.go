package main

import (
	"fmt"
	"os"
	"strconv"
	"time"
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

	results := []Result{}
	for t := 0; t < tryTimes; t++ {
		result := playPreFlop(hands)
		results = append(results, result)

	}

	resultCount = calcResultCount(results)

	fmt.Printf("OnePair        : %d ( %f %v )\n", resultCount.CountOnePair, float64(resultCount.CountOnePair)/float64(tryTimes)*100, "%")
	fmt.Printf("TwoPair        : %d ( %f %v )\n", resultCount.CountTwoPair, float64(resultCount.CountTwoPair)/float64(tryTimes)*100, "%")
	fmt.Printf("ThreeOfAKind   : %d ( %f %v )\n", resultCount.CountThreeOfAKind, float64(resultCount.CountThreeOfAKind)/float64(tryTimes)*100, "%")
	fmt.Printf("Straight       : %d ( %f %v )\n", resultCount.CountStraight, float64(resultCount.CountStraight)/float64(tryTimes)*100, "%")
	fmt.Printf("Flush          : %d ( %f %v )\n", resultCount.CountFlush, float64(resultCount.CountFlush)/float64(tryTimes)*100, "%")
	fmt.Printf("FullHouse      : %d ( %f %v )\n", resultCount.CountFullHouse, float64(resultCount.CountFullHouse)/float64(tryTimes)*100, "%")
	fmt.Printf("FourOfAKind    : %d ( %f %v )\n", resultCount.CountFourOfAKind, float64(resultCount.CountFourOfAKind)/float64(tryTimes)*100, "%")
	fmt.Printf("StraightFlush  : %d ( %f %v )\n", resultCount.CountStraightFlush, float64(resultCount.CountStraightFlush)/float64(tryTimes)*100, "%")

	elapsed := time.Since(start)
	fmt.Println(elapsed)
}
