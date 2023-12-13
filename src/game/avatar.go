package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
	"image/color"
)

func getAvatarPos(game *Game, screen *ebiten.Image) (float64, float64) {
	w, h := getImageSize(screen)
	x := game.x + float64(w/2)
	y := game.y + float64(h/2)
	return x, y
}

func DrawAvatar(screen *ebiten.Image, g *Game) {
	teal := color.RGBA{R: 0, G: 80, B: 80, A: 1}
	green := color.RGBA{R: 0, G: 0xFF, B: 80, A: 1}
	miniSquareSize := float64(SquareSize / 3)

	x, y := getAvatarPos(g, screen)
	vector.DrawFilledRect(screen, float32(x-miniSquareSize/2), float32(y-miniSquareSize/2), float32(miniSquareSize), float32(miniSquareSize), teal, true)
	vector.DrawFilledCircle(screen, float32(x+SquareSize*2), float32(y), SquareSize, green, true)
	vector.DrawFilledRect(screen, float32(x-SquareSize/2), float32(y-SquareSize/2), SquareSize, SquareSize, teal, true)
	vector.DrawFilledRect(screen, float32(x+2*SquareSize-SquareSize/2), float32(y+2*SquareSize-SquareSize/2), SquareSize, SquareSize, teal, true)
}
