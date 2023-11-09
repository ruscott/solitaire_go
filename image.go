package main

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"
)

func drawImages(screen *ebiten.Image, imageKey string) {
	suite, ok := GetImage(imageKey)
	if !ok {
		log.Fatal("Failed to retrieve the '%s' image.", imageKey)
		return
	}

	optionsLeft := &ebiten.DrawImageOptions{}
	screen.DrawImage(suite, optionsLeft)

	optionsRight := &ebiten.DrawImageOptions{}
	optionsRight.GeoM.Translate(50, 50)
	screen.DrawImage(suite, optionsRight)
}