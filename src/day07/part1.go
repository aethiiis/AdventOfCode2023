package main

import (
	"advent_of_code_2023/src/utilities"
	"fmt"
	"sort"
	"strconv"
	"strings"
	"unicode"
)

type HandsPart1 []utilities.Hand

func (h HandsPart1) Len() int {
	return len(h)
}

func (h HandsPart1) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h HandsPart1) Less(i, j int) bool {
	return !h[i].Compare(h[j])
}

func Part1(lines []string) int {
	var tab HandsPart1
	tab = GetTabPart1(lines)
	sort.Sort(tab)

	winnings := 0
	for i := range tab {
		winnings += tab[i].Bid * (i + 1)
	}
	return winnings
}

func GetTabPart1(lines []string) []utilities.Hand {
	var tab []utilities.Hand

	for i := range lines {
		cardsBid := strings.SplitN(lines[i], " ", 2)
		cardsString := cardsBid[0]
		bidString := cardsBid[1]

		var force int
		var cards []int
		var bid int

		countCards := make([]int, 13)

		// Cards
		for j := range cardsString {
			if unicode.IsDigit(rune(cardsString[j])) {
				cards = append(cards, int(cardsString[j]-48))
				countCards[int(cardsString[j]-48)-2]++
			} else if rune(cardsString[j]) == 'T' {
				cards = append(cards, 10)
				countCards[8]++
			} else if rune(cardsString[j]) == 'J' {
				cards = append(cards, 11)
				countCards[9]++
			} else if rune(cardsString[j]) == 'Q' {
				cards = append(cards, 12)
				countCards[10]++
			} else if rune(cardsString[j]) == 'K' {
				cards = append(cards, 13)
				countCards[11]++
			} else if rune(cardsString[j]) == 'A' {
				cards = append(cards, 14)
				countCards[12]++
			}
		}

		// Force
		if utilities.IsIn(countCards, 5) {
			force = 7
		} else if utilities.IsIn(countCards, 4) {
			force = 6
		} else if utilities.IsIn(countCards, 3) && utilities.IsIn(countCards, 2) {
			force = 5
		} else if utilities.IsIn(countCards, 3) {
			force = 4
		} else if utilities.IsInTwo(countCards, 2) {
			force = 3
		} else if utilities.IsIn(countCards, 2) {
			force = 2
		} else if utilities.IsIn(countCards, 1) {
			force = 1
		} else {
			panic("Main non-reconnue")
		}

		// Bid
		chiffre, err := strconv.Atoi(bidString)
		if err != nil {
			fmt.Println("Erreur de conversion string --> int", err)
		} else {
			bid = chiffre
		}

		hand := utilities.Hand{Force: force, Cards: cards, Bid: bid}
		tab = append(tab, hand)
	}
	return tab
}
