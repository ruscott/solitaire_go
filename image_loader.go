package main

import (
	"log"
	"sync"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

var (
	once        sync.Once
	images      map[string]*ebiten.Image
	imagesReady bool
)

func initImages() {
	images = make(map[string]*ebiten.Image)
	loadImage("hearts", "images/suite/hearts.png")
	loadImage("clubs", "images/suite/clubs.png")
	loadImage("spades", "images/suite/spades.png")
	loadImage("diamonds", "images/suite/diamonds.png")
	loadImage("card", "images/card.png")
	imagesReady = true
}

func loadImage(key, path string) {
	img, _, err := ebitenutil.NewImageFromFile(path)
	if err != nil {
		log.Fatal(err)
	}
	images[key] = img
}

func GetImage(key string) (*ebiten.Image, bool) {
	if !imagesReady {
		log.Fatal("Images not initialized. Call initImages first.")
		return nil, false
	}

	img, ok := images[key]
	return img, ok
}