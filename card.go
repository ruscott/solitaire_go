package main

import "fmt"

type Card struct {
	Suite int
	Value int
}

type CardConfig struct {
	Suites []int
	Values []int
}

var cardConfig = CardConfig{
	Suites: []int{1,2,3,4},
	Values: []int{2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 1},
}

func (card Card) String() string {
	return fmt.Sprintf("%d of %d", card.Value, card.Suite)
}

func rankToString(value int) string {
	ranks := map[int]string{
		2: "Two",
		3: "Three",
		4: "Four",
		5: "Five",
		6: "Six",
		7: "Seven",
		8: "Eight",
		9: "Nine",
		10: "Ten",
		11: "Jack",
		12: "Queen",
		13: "King",
		1: "Ace",
	}

	return ranks[value]
}

func suiteToString(value int) string {
	ranks := map[int]string{
		1: "Spades",
		2: "Clubs",
		3: "Diamonds",
		4: "Hearts",
	}

	return ranks[value]
}

func getColor(card Card) string {
	switch card.Suite {
	case 3, 4:
		return "RED"
	case 1, 2:
		return "BLACK"
	default:
		return "NONE"
	}
}