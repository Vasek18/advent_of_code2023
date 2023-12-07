package main

import (
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Hand struct {
	hand                string
	combinationStrength float32
	cardsStrength       int
	bid                 int
}

func solve(hands []Hand) int {
	sum := 0

	sort.Slice(hands, func(i, j int) bool {
		hand1 := hands[i]
		hand2 := hands[j]

		if hand1.combinationStrength > hand2.combinationStrength {
			return false
		}
		if hand1.combinationStrength < hand2.combinationStrength {
			return true
		}
		if hand1.cardsStrength > hand2.cardsStrength {
			return false
		}
		if hand1.cardsStrength < hand2.cardsStrength {
			return true
		}

		return false
	})

	// fmt.Println(hands)

	for i, hand := range hands {
		fmt.Println((i + 1), hand.bid, hand.hand)
		sum += (i + 1) * hand.bid
	}

	return sum
}

func getCombinationStrength(hand string) float32 {
	charCount := make(map[rune]int)

	for _, char := range hand {
		charCount[char]++
	}

	numberOfPairs := 0
	numberOfThrees := 0
	for _, count := range charCount {
		if count > 3 {
			return float32(count)
		}

		if count == 2 {
			numberOfPairs++
		}
		if count == 3 {
			numberOfThrees++
		}
	}

	if numberOfPairs == 1 && numberOfThrees == 1 {
		return 3.5
	}

	if numberOfThrees == 1 {
		return 3
	}

	if numberOfPairs == 2 {
		return 2.5
	}

	if numberOfPairs == 1 {
		return 2
	}

	return 1
}

func getCardsStrength(hand string) int {
	strength := 0

	for i, card := range hand {
		multiplier := int(math.Pow(10, float64(5-i)))

		if string(card) == "A" {
			strength += 14 * multiplier
			continue
		}
		if string(card) == "K" {
			strength += 13 * multiplier
			continue
		}
		if string(card) == "Q" {
			strength += 12 * multiplier
			continue
		}
		if string(card) == "J" {
			strength += 11 * multiplier
			continue
		}
		if string(card) == "T" {
			strength += 10 * multiplier
			continue
		}

		strength += convertToInt(string(card)) * multiplier
	}

	return strength
}

func prepareHands(inputStrings []string) []Hand {
	var hands []Hand

	for _, rawHand := range inputStrings {

		rawHandParts := strings.Fields(rawHand)
		hand := Hand{
			hand:                rawHandParts[0],
			combinationStrength: getCombinationStrength(rawHandParts[0]),
			cardsStrength:       getCardsStrength(rawHandParts[0]),
			bid:                 convertToInt(rawHandParts[1]),
		}

		hands = append(hands, hand)
	}

	return hands
}

func main() {
	input, _ := os.ReadFile("./07.1/input.txt")

	hands := prepareHands(strings.Split(string(input), "\n"))

	totalSum := solve(hands)

	fmt.Println(totalSum)
}

func convertToInt(s string) int {
	n, _ := strconv.Atoi(s)
	return n
}
