package main

import (
	"fmt"
	"os"
	"strconv"
)

func splitSuitsAndNumbers(chars string) ([]string, []string) {
	s := []string{"s", "h", "d", "c"}
	suits := []string{}
	nums := []string{}
	isSuit := false
	for _, char := range chars {
		c := string(char)
		if containsString(s, c) {
			suits = append(suits, c)
			isSuit = true
		} else {
			if isSuit {
				nums = append(nums, c)
			} else {
				nums[len(nums)-1] = nums[len(nums)-1] + c
			}
			isSuit = false
		}
	}
	return suits, nums
}

func createCardsFromSuitsAndNumbers(suits []string, numbers []string) []Card {
	cards := []Card{}
	for i, s := range suits {
		n, _ := strconv.Atoi(numbers[i])
		card := Card{s, n}
		if !card.Valid() {
			fmt.Println("invalid card")
			os.Exit(1)
		}
		cards = append(cards, card)
	}
	return cards
}

func containsString(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

func splitTryTimes(tryTimes int) []int {
	times := []int{}
	timesCount := tryTimes / 10000
	for t := 0; t < timesCount; t++ {
		times = append(times, 10000)
	}

	rest := tryTimes % 10000
	if rest > 0 {
		times = append(times, rest)
	}
	return times
}
