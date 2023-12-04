package main

import (
	"github.com/hajimehoshi/ebiten/v2"
)

func drawImages(screen *ebiten.Image, imageKey string, x, y int) {
    suite, _ := GetImage(imageKey)

    // Draw the image on the screen
    op := &ebiten.DrawImageOptions{}
    op.GeoM.Translate(float64(x), float64(y))
    screen.DrawImage(suite, op)
}