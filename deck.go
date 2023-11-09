package main

import (
	"math/rand"
)

type Deck struct {
	Cards []Card
}

type Card struct {
	Suite string
	Value string
}

func shuffleDeck(deck Deck) Deck {
	cards := deck.Cards
	for i := range cards {
		j := rand.Intn(i + 1)
		cards[i], cards[j] = cards[j], cards[i]
	}
	return deck
}

func makeDeck() Deck {
	suites := []string{"spades", "hearts", "diamonds", "clubs"}
	values := []string{"two", "three", "four", "five", "six", "seven", "eight", "nine", "ten", "Jack", "Queen", "King", "Ace"}
	deck := Deck{}
	for _, suite := range suites {
		for _, value := range values {
			deck.Cards = append(deck.Cards, Card{Suite: suite, Value: value})
		}
	}
	return deck
}

func getFirstCardFromShuffledDeck() Card {
	deck := makeDeck()
	shuffledDeck := shuffleDeck(deck)
	return shuffledDeck.Cards[0]
}