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
	NoCard = Card{}

)

type Game struct {
	Deck         []Card
	CurrentCard  Card
	Refresh      bool
	refreshCount int
	SolitaireSet SolitaireSet

	GameState      int
	SelectedColumn int
	SelectedCard Card

}
func (g *Game) Init() {
	initImages()
	if g.CurrentCard == (Card{}) {
		g.CurrentCard = getFirstCardFromShuffledDeck()
	}

	if len(g.SolitaireSet.CardSets) == 0 {
		deck := makeDeck()
		shuffledDeck := shuffleDeck(deck)
		g.SolitaireSet = createSolitaireSet(shuffledDeck)
	}

	g.GameState = StateSelectCard
	g.SelectedColumn = NoSelection
	g.SelectedCard = NoCard
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
		if isKeyJustPressed(ebiten.KeyD) {
			if len(g.SolitaireSet.CardSets[setNumber]) > 0 {
				card := g.SolitaireSet.CardSets[setNumber][0]
				g.SolitaireSet.CardSets[setNumber] = append(g.SolitaireSet.CardSets[setNumber][1:], card)
			}
		}
		for i := 0; i <= 9; i++ {
			key := ebiten.Key(i) + ebiten.Key0
			if isKeyJustPressed(key) {
				columnNumber := i
				if columnNumber == 0{
					columnNumber = setNumber+1} else if columnNumber >= 0 && columnNumber < len(g.SolitaireSet.CardSets) && len(g.SolitaireSet.CardSets[columnNumber-1]) > 0 {
						g.SelectedCard = g.SolitaireSet.CardSets[columnNumber-1][0]
					}
					g.GameState = StateSelectDestination
					g.SelectedColumn = columnNumber - 1
				}
			}

	case StateSelectDestination:
		for i := 0; i <= 9; i++ {
			key := ebiten.Key(i) + ebiten.Key0
			if isKeyJustPressed(key) {
				destinationKey := i
				if destinationKey > 0 && destinationKey < len(g.SolitaireSet.CardSets)-1 {
					if len(g.SolitaireSet.CardSets[destinationKey-1]) > 0 {
						if isValidMove(g.SelectedCard, g.SolitaireSet.CardSets[destinationKey-1]) {
							g.SolitaireSet.CardSets[destinationKey-1] = append([]Card{g.SelectedCard}, g.SolitaireSet.CardSets[destinationKey-1]...)
							g.SolitaireSet.CardSets[g.SelectedColumn] = g.SolitaireSet.CardSets[g.SelectedColumn][1:]

						}
					}
					g.GameState = StateSelectCard
					g.SelectedCard = NoCard
				}
			}
		}
	}

	return nil
}

func isValidMove(card Card, destinationColumn []Card) bool {
	if len(destinationColumn) == 0 {
		fmt.Printf("card %d", card.Value)
		return card.Value == 13
	}

	topCard := destinationColumn[0]
	topCardColour := getColor(topCard)
	cardColour := getColor(card)

	fmt.Printf("card 1 %d, card2 %d", topCard.Value, card.Value)
	return topCard.Value == card.Value+1 && topCardColour != cardColour
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(backgroundColor)
	cellWidth := float64(screen.Bounds().Dx()) / float64(len(g.SolitaireSet.CardSets))

	drawPrompt(screen, g.GameState)

	for i, cardSet := range g.SolitaireSet.CardSets {
		if i == setNumber {
			drawCardInfo(screen, cardSet, ScreenWidth-40, -1, 20, 20)
		}else{
		x := int(float64(i) * cellWidth)
		drawCardInfo(screen, cardSet, x, i, 60, 20)
		}
	}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return ScreenWidth, ScreenHeight
}
