package main

import (
	"fmt"
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
)

const (
	StateSelectCard = iota
	StateSelectDestination
	NoSelection = -1
)

var (
	backgroundColor = color.RGBA{39, 92, 39, 1}
)

type Game struct {
	Deck         []Card
	CurrentCard  Card
	Refresh      bool
	refreshCount int
	SolitaireSet SolitaireSet

	GameState      int
	SelectedColumn int

}
func (g *Game) Init() {
	initImages()
	if g.CurrentCard == (Card{}) {
		g.CurrentCard = getFirstCardFromShuffledDeck()
	}

	if len(g.SolitaireSet.CardSets) == 0 {
		deck := makeDeck()
		shuffledDeck := shuffleDeck(deck)
		g.SolitaireSet = createSolitaireSets(shuffledDeck)
	}

	g.GameState = StateSelectCard
	g.SelectedColumn = NoSelection
}

var keyPressedMap = make(map[ebiten.Key]bool)

func isKeyJustPressed(key ebiten.Key) bool {
	if ebiten.IsKeyPressed(key) && !keyPressedMap[key] {
		keyPressedMap[key] = true
		return true
	} else if !ebiten.IsKeyPressed(key) {
		keyPressedMap[key] = false
	}
	return false
}
func (g *Game) Update() error {
	switch g.GameState {
	case StateSelectCard:
		for i := 0; i <= 9; i++ {
			key := ebiten.Key(i) + ebiten.Key0
			if isKeyJustPressed(key) {
				columnNumber := i
				fmt.Printf("column length %d ", len(g.SolitaireSet.CardSets[columnNumber-1]))
				if columnNumber >= 0 && columnNumber < len(g.SolitaireSet.CardSets) && len(g.SolitaireSet.CardSets[columnNumber-1]) > 0 {
					g.SelectedColumn = columnNumber
					g.GameState = StateSelectDestination
				}
			}
		}
	case StateSelectDestination:
		for i := 0; i <= 9; i++ {
			key := ebiten.Key(i) + ebiten.Key0
			if isKeyJustPressed(key) {
				destinationKey := i
				fmt.Printf("dest no %d ", destinationKey)

				if destinationKey >= 0 && destinationKey < len(g.SolitaireSet.CardSets) {
					if len(g.SolitaireSet.CardSets[g.SelectedColumn]) > 0 {
						card := g.SolitaireSet.CardSets[g.SelectedColumn-1][0]

						if isValidMove(card, g.SolitaireSet.CardSets[destinationKey-1]) {
							g.SolitaireSet.CardSets[g.SelectedColumn-1] = g.SolitaireSet.CardSets[g.SelectedColumn-1][1:]
							g.SolitaireSet.CardSets[destinationKey-1] = append([]Card{card}, g.SolitaireSet.CardSets[destinationKey-1]...)
						}
					}

					g.SelectedColumn = NoSelection
					g.GameState = StateSelectCard
				}
			}
		}
	}

	return nil
}

func isValidMove(card Card, destinationColumn []Card) bool {
	if len(destinationColumn) == 0 {
		return card.Value == 13
	}

	topCard := destinationColumn[0]
	topCardColour := getColor(topCard)
	cardColour := getColor(card)
	return topCard.Value == card.Value+1 && topCardColour != cardColour
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(backgroundColor)
	cellWidth := float64(screen.Bounds().Dx()) / float64(len(g.SolitaireSet.CardSets))

	drawPrompt(screen, g.GameState)

	for i, cardSet := range g.SolitaireSet.CardSets {
		x := int(float64(i) * cellWidth)
		drawCardInfo(screen, cardSet, x, i, 20, 20)
	}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return ScreenWidth, ScreenHeight
}
