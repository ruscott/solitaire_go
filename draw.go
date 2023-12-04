package main

import (
	"fmt"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

func drawPrompt(screen *ebiten.Image, state int) {
	var prompt string
	switch state {
	case StateSelectCard:
		prompt = "Please enter the number of the card to move (0-9):"
	case StateSelectDestination:
		prompt = "Please enter the destination number (0-9):"
	}
	ebitenutil.DebugPrintAt(screen, prompt, 10, screen.Bounds().Dy()-20)
}

func drawCardInfo(screen *ebiten.Image, cardSet []Card, x, yCardInfo, yOf, ySuit int) {
	drawInfo := func(info string, y int) {
		ebitenutil.DebugPrintAt(screen, info, x, y)
	}

	if len(cardSet) > 0 {
		firstCard := cardSet[0]
		drawInfo(fmt.Sprintf("%s", firstCard.Value), yCardInfo)
		drawInfo("of", yOf)
		drawInfo(fmt.Sprintf("%s", firstCard.Suite), ySuit)
	} else {
		drawInfo("No", yCardInfo)
		drawInfo("Value", yOf)
	}
}

func drawCardImages(screen *ebiten.Image, cardSets [][]Card, cellWidth float64) {
	for i, cardSet := range cardSets {
		if len(cardSet) > 0 {
			firstCard := cardSet[0]
			x := int(float64(i) * cellWidth)
			drawImages(screen, firstCard.Suite, x, 100)
		}
	}
}