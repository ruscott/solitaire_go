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

func drawCardInfo(screen *ebiten.Image, cardSet []Card, x, columnNumber, yStart, spacing int) {
	drawInfo := func(info string, y int) {
		ebitenutil.DebugPrintAt(screen, info, x, y)
	}
	
	y := yStart

	if len(cardSet) > 0 {
		firstCard := cardSet[0]
		drawInfo(fmt.Sprintf("%d",columnNumber+1), y)
		y = y+spacing
		drawInfo(fmt.Sprintf("%s", rankToString(firstCard.Value)), y)
		y=y+spacing
		drawInfo("of", y)
		y=y+spacing
		drawInfo(suiteToString(firstCard.Suite), y)
		y=y+spacing
		drawCardImages(screen, firstCard, x, y)
	} else {
		drawInfo(fmt.Sprintf("%d",columnNumber+1), y)
		y=y+spacing
		drawInfo("No", y)
		y=y+spacing
		drawInfo("Value", y)
	}
}

func drawCardImages(screen *ebiten.Image, card Card, cellWidth int, y int) {
	drawImages(screen, suiteToString(card.Suite), int(cellWidth), y)
		}