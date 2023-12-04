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
	SourceColumn   int
	KeyPressed    bool

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
	g.SourceColumn = NoSelection
	g.KeyPressed = false
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
	g.KeyPressed = false
	switch g.GameState {
	case StateSelectCard:
		for i := 0; i <= 9; i++ {
			key := ebiten.Key(i) + ebiten.Key0
			if isKeyJustPressed(key) && !g.KeyPressed {
				columnNumber := i
				fmt.Printf("column no %d ", columnNumber)
				if columnNumber >= 0 && columnNumber < len(g.SolitaireSet.CardSets) && len(g.SolitaireSet.CardSets[columnNumber]) > 0 {
					g.SourceColumn = columnNumber
					g.GameState = StateSelectDestination
					g.KeyPressed = true
				}
			}
		}
	case StateSelectDestination:
		for i := 0; i <= 9; i++ {
			key := ebiten.Key(i) + ebiten.Key0
			if isKeyJustPressed(key) && !g.KeyPressed {
				destinationKey := i
				fmt.Printf("dest no %d ", destinationKey)

				if destinationKey >= 0 && destinationKey < len(g.SolitaireSet.CardSets) {
					if len(g.SolitaireSet.CardSets[g.SourceColumn]) > 0 {
						card := g.SolitaireSet.CardSets[g.SourceColumn-1][0]
						g.SolitaireSet.CardSets[g.SourceColumn-1] = g.SolitaireSet.CardSets[g.SourceColumn-1][1:]
						fmt.Printf("First card in destination column %d: %+v\n", destinationKey-1, g.SolitaireSet.CardSets[destinationKey-1][0])

						g.SolitaireSet.CardSets[destinationKey-1] = append([]Card{card}, g.SolitaireSet.CardSets[destinationKey-1]...)
					}

					g.SelectedColumn = NoSelection 
					g.SourceColumn = NoSelection
					g.GameState = StateSelectCard
					g.KeyPressed = true
				}
			}
		}
	}

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(backgroundColor)
	cellWidth := float64(screen.Bounds().Dx()) / float64(len(g.SolitaireSet.CardSets))

	drawPrompt(screen, g.GameState)

	for i, cardSet := range g.SolitaireSet.CardSets {
		x := int(float64(i) * cellWidth)
		yCardInfo, ySuit, yOf := 40, 80, 60
		drawCardInfo(screen, cardSet, x, yCardInfo, yOf, ySuit)
	}

	drawCardImages(screen, g.SolitaireSet.CardSets, cellWidth)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 640, 240
}