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
	loadImage("Hearts", "images/suite/hearts.png")
	loadImage("Clubs", "images/suite/clubs.png")
	loadImage("Spades", "images/suite/spades.png")
	loadImage("Diamonds", "images/suite/diamonds.png")
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