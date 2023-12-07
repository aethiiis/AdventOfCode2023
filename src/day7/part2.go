package main

import (
	"advent_of_code_2023/src/utilities"
	"fmt"
	"sort"
	"strconv"
	"strings"
	"unicode"
)

type HandsPart2 []utilities.Hand

func (h HandsPart2) Len() int {
	return len(h)
}

func (h HandsPart2) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h HandsPart2) Less(i, j int) bool {
	return !h[i].Compare(h[j])
}

func Part2(lines []string) int {
	var tab HandsPart2
	tab = GetTabPart2(lines)
	sort.Sort(tab)

	winnings := 0
	for i := range tab {
		winnings += tab[i].Bid * (i + 1)
	}
	return winnings
}

func GetTabPart2(lines []string) []utilities.Hand {
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
			if rune(cardsString[j]) == 'J' {
				cards = append(cards, 0)
				countCards[9]++
			} else if unicode.IsDigit(rune(cardsString[j])) {
				cards = append(cards, int(cardsString[j]-48))
				countCards[int(cardsString[j]-48)-2]++
			} else if rune(cardsString[j]) == 'T' {
				cards = append(cards, 10)
				countCards[8]++
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

		// Cas Joker(s)
		if countCards[9] != 0 {
			if force == 7 {
				// JJJJJ
			} else if force == 6 {
				// JJJJA ou JAAAA
				force += 1
			} else if force == 5 {
				// JJJAA ou JJAAA
				force += 2
			} else if force == 4 {
				// JJJAK ou JAKKK
				force += 2
			} else if force == 3 {
				if countCards[9] == 2 {
					// JJAAK
					force += 3
				} else if countCards[9] == 1 {
					// JAAKK
					force += 2
				} else {
					panic("Erreur dans le comptage des Jokers : 1")
				}
			} else if force == 2 {
				// JJAKQ ou JAAKQ
				force += 2
			} else if force == 1 {
				// JAKQT
				force += 1
			} else {
				panic("Erreur dans le comptage des Jokers : 2")
			}
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
