package main

import (
	"log"

	"game/internal/assets"
	"game/internal/game"

	"github.com/hajimehoshi/ebiten/v2"
)

func main() {
	scaleFactor := 0.03125

	if err := assets.InitAssetsWithScale(scaleFactor); err != nil {
		log.Printf("Warning: Could not load all assets: %v", err)
		log.Println("Using fallback colored rectangles")
	}

	gameInstance := game.NewGame()

	ebiten.SetWindowSize(game.ScreenWidth, game.ScreenHeight)
	ebiten.SetWindowTitle("Hero Battle")

	if err := ebiten.RunGame(gameInstance); err != nil {
		log.Fatal(err)
	}
}
