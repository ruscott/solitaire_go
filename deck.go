package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Deck struct {
	Cards []Card
}


type SolitaireSet struct {
	CardSets [][]Card
}

func shuffleDeck(deck Deck) Deck {
	rand.Seed(time.Now().UnixNano())
	cards := deck.Cards
	for i := range cards {
		j := rand.Intn(i + 1)
		cards[i], cards[j] = cards[j], cards[i]
	}
	return deck
}

func makeDeck() Deck {
	deck := Deck{}
	for _, suite := range cardConfig.Suites {
		for _, value := range cardConfig.Values {
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

func createSolitaireSets(deck Deck) SolitaireSet {
	shuffledDeck := shuffleDeck(deck)
	solitaireSet := SolitaireSet{}

	cardLen := 1
	iteration := 1
	for cardLen < len(shuffledDeck.Cards) {
		deckSlice := shuffledDeck.Cards[cardLen : iteration+cardLen]
		cardLen = cardLen + iteration
		iteration = iteration + 1
		solitaireSet.CardSets = append(solitaireSet.CardSets, deckSlice)
	}

	return solitaireSet
}

func printFirstCardOfSets(solitaireSet SolitaireSet) {
	// Determine the number of sets
	numSets := len(solitaireSet.CardSets)

	// Print set numbers in the top row
	for i := 0; i < numSets; i++ {
		fmt.Printf("%d\t", i+1)
	}
	fmt.Println()

	// Print the first card of each set in the bottom row
	for _, cardSet := range solitaireSet.CardSets {
		if len(cardSet) > 0 {
			firstCard := cardSet[0]
			fmt.Printf("%d of %d\t", firstCard.Value, firstCard.Suite)
		} else {
			fmt.Print("Empty\t")
		}
	}
	fmt.Println()
}





