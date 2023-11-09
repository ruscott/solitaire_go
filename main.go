package main

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"
)

var (
	image *ebiten.Image
)

func main() {
	ebiten.SetWindowSize(640, 480)
	ebiten.SetWindowTitle("Card Game")
	if err := ebiten.RunGame(&Game{}); err != nil {
		log.Fatal(err)
	}
}