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
	CurrentCard Card
}

func (g *Game) Update() error {
	if g.CurrentCard == (Card{}) {
		g.CurrentCard = getFirstCardFromShuffledDeck()
	}
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(backgroundColor)

	once.Do(initImages)

	cardInfo := fmt.Sprintf("Current Card: %s of %s", g.CurrentCard.Value, g.CurrentCard.Suite)
	ebitenutil.DebugPrint(screen, cardInfo)

	drawImages(screen, g.CurrentCard.Suite)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 320, 240
}
