package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

type hand struct {
	cards string
	bid   int64
}

// A, K, Q, J, T, 9, 8, 7, 6, 5, 4, 3, or 2
func getScore(card byte) int {
	rank := "J23456789TQKA"
	return strings.Index(rank, string(card))
}

func getRealRank(cards1 string) int {
	rankList := "23456789TJQKA"
	rank := 0
	for _, cardToBecome := range rankList {
		newCard := strings.Replace(cards1, "J", string(cardToBecome), -1)
		newRank := getRank(newCard)
		if newRank > rank {
			rank = newRank
		}
	}
	return rank
}

// High card: 1, One pair: 2, Two pairs: 3, Three of a kind: 4, Full house: 5, Four of a kind: 6, Five of a kind: 7
func getRank(cards1 string) int {
	cards := make(map[string]int)
	numberOfJoker := 0
	for _, card := range cards1 {
		cards[string(card)]++
		if string(card) == "J" {
			numberOfJoker++
		}
	}

	if len(cards) == 5 {
		return 1
	} else if len(cards) == 4 {
		return 2
	} else if len(cards) == 3 {
		isTwoPair := false
		for _, v := range cards {
			if v == 2 {
				isTwoPair = true
			}
		}
		if isTwoPair {
			return 3
		} else {
			return 4
		}
	} else if len(cards) == 2 {
		isFullHouse := false
		for _, v := range cards {
			if v == 3 {
				isFullHouse = true
			}
		}
		if isFullHouse {
			return 5
		} else {
			return 6
		}
	} else {
		return 7
	}
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	hands := make([]hand, 0)

	for scanner.Scan() {
		var cards string
		var bid int64
		fmt.Sscanf(scanner.Text(), "%s %d", &cards, &bid)
		hands = append(hands, hand{cards, bid})
		// fmt.Println(cards, getScore(cards[0]))
	}

	sort.Slice(hands, func(i, j int) bool {
		if getRealRank(hands[i].cards) == getRealRank(hands[j].cards) {
			for k := 0; k < 5; k++ {
				if getScore(hands[i].cards[k]) != getScore(hands[j].cards[k]) {
					return getScore(hands[i].cards[k]) < getScore(hands[j].cards[k])
				}
			}
		}

		return getRealRank(hands[i].cards) < getRealRank(hands[j].cards)
	})

	result := int64(0)
	for i, hand := range hands {
		result += int64(i+1) * hand.bid
	}

	fmt.Println(result)
}
