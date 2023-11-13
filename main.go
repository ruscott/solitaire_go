package main

import (
	"github.com/hajimehoshi/ebiten/v2"
)

var (
	imageEb *ebiten.Image
)

func main() {
	ebiten.SetWindowSize(640, 480)
	ebiten.SetWindowTitle("Card Game")
	
	game := &Game{}
	game.Init()

	ebiten.RunGame(game)
}
	