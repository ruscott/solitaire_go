package main

import (
	"fmt"
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
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
}

func (g *Game) Init() {
	if g.CurrentCard == (Card{}) {
		g.CurrentCard = getFirstCardFromShuffledDeck()
	}

	if len(g.SolitaireSet.CardSets) == 0 {
		deck := makeDeck()
		shuffledDeck := shuffleDeck(deck)
		g.SolitaireSet = createSolitaireSets(shuffledDeck)
	}
}

func (g *Game) Update() error {
	return nil
}


func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(backgroundColor)
	cellWidth := float64(screen.Bounds().Dx()) / float64(len(g.SolitaireSet.CardSets))

	for i := range g.SolitaireSet.CardSets {
		setNumber := fmt.Sprintf("%d", i+1)
		x := float64(i) * cellWidth
		ebitenutil.DebugPrintAt(screen, setNumber, int(x), 20)
	}

	for i, cardSet := range g.SolitaireSet.CardSets {
		if len(cardSet) > 0 {
			firstCard := cardSet[0]
			cardInfo := fmt.Sprintf("%s", firstCard.Value)
			x := float64(i) * cellWidth
			ebitenutil.DebugPrintAt(screen, cardInfo, int(x), 40)
		} else {
			x := float64(i) * cellWidth
			ebitenutil.DebugPrintAt(screen, "No", int(x), 40)
		}
	}

	for i, cardSet := range g.SolitaireSet.CardSets {
		if len(cardSet) > 0 {
			firstCard := cardSet[0]
			cardInfo := fmt.Sprintf("%s", firstCard.Suite)
			x := float64(i) * cellWidth
			ebitenutil.DebugPrintAt(screen, cardInfo, int(x), 80)
		} else {
			x := float64(i) * cellWidth
			ebitenutil.DebugPrintAt(screen, "Value", int(x), 80)
		}
	}


	for i, cardSet := range g.SolitaireSet.CardSets {
		if len(cardSet) > 0 {
			cardInfo := fmt.Sprintf("of")
			x := float64(i) * cellWidth
			ebitenutil.DebugPrintAt(screen, cardInfo, int(x), 60)
		} else {
			x := float64(i) * cellWidth
			ebitenutil.DebugPrintAt(screen, "", int(x), 60)
		}
	}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 640, 240
}

func (g *Game) drawSolitaireSet(screen *ebiten.Image) {
	for i, cardSet := range g.SolitaireSet.CardSets {
		if len(cardSet) > 0 {
			firstCard := cardSet[0]
			info := fmt.Sprintf("Set %d: %s of %s", i+1, firstCard.Value, firstCard.Suite)
			ebitenutil.DebugPrint(screen, info)
		} else {
			info := fmt.Sprintf("Set %d is empty", i+1)
			ebitenutil.DebugPrint(screen, info)
		}
	}
}