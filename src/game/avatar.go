package main

import (
	"fmt"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text"
	"github.com/hajimehoshi/ebiten/v2/vector"
	"image/color"
	"math"
)

func getAvatarPos(game *Game, screen *ebiten.Image) (float64, float64) {
	w, _ := getImageSize(screen)
	multX := 0
	//multY := 1 / 2
	x := game.x + float64(w*multX)
	y := game.y + 150
	return x, y
}

const SquareSize = 50

func DrawAvatar(screen *ebiten.Image, g *Game) {
	size := SquareSize * g.zoom
	x, y := getAvatarPos(g, screen)
	drawSquares(screen, x, y, size)
	drawSeekers(screen, g.x, g.zoom)
}

func drawSquares(screen *ebiten.Image, x float64, y float64, size float64) {
	teal := color.RGBA{R: 0, G: 80, B: 80, A: 1}
	//green := color.RGBA{R: 0, G: 0xFF, B: 80, A: 1}
	miniSquareSize := float64(size / 3)
	vector.DrawFilledRect(screen, float32(x-miniSquareSize/2), float32(y-miniSquareSize/2), float32(miniSquareSize), float32(miniSquareSize), teal, true)
	text.Draw(screen, fmt.Sprintf("%d", int(x)), standardFontFace, int(x), int(y), color.White)
	//vector.DrawFilledRect(screen, float32(x-size/2), float32(y-size/2), float32(size), float32(size), teal, true)
	//vector.DrawFilledCircle(screen, float32(x+size*2), float32(y), float32(size), teal, true)
	//vector.DrawFilledRect(screen, float32(x+2*size-size/2), float32(y+2*size-size/2), float32(size), float32(size), teal, true)
}

type Marker struct {
	value float64
	color color.RGBA
	name  string
}

func drawSeekers(screen *ebiten.Image, counter float64, scale float64) {
	factor := 10
	k := math.Abs(scale / float64(factor))
	_, h := getImageSize(screen)
	h -= 50
	size := 100
	xVal := counter * k
	text.Draw(screen, fmt.Sprintf("scale: %.3f/%.3f", scale, float32(factor)), standardFontFace, 0, 20, color.RGBA{G: 255, A: 255})
	text.Draw(screen, fmt.Sprintf("scale: %.3f", k), standardFontFace, 0, 50, color.RGBA{G: 255, A: 255})
	text.Draw(screen, fmt.Sprintf("n: %.2f", xVal), standardFontFace, 0, 100, color.RGBA{G: 255, A: 255})
	// squares at the bottom of the screen

	y0 := float32(h - size)

	shapes := []Marker{
		{name: "n!", value: math.Gamma(xVal + 1), color: color.RGBA{R: 80, G: 0, B: 80, A: 1}},
		{name: "2^n", value: math.Pow(2, xVal), color: color.RGBA{R: 0, G: 0, B: 80, A: 1}},
		{name: "n^2", value: math.Pow(xVal, 2), color: color.RGBA{R: 80, G: 0, B: 0, A: 1}},
		{name: "n log(n)", value: xVal * math.Log(xVal), color: color.RGBA{R: 0, G: 80, B: 80, A: 1}},
		{name: "n", value: math.Pow(xVal, 1), color: color.RGBA{R: 80, G: 80, B: 0, A: 1}},
		{name: "log(n)", value: math.Log(xVal), color: color.RGBA{R: 0, G: 80, B: 0, A: 1}},
	}

	for i, shape := range shapes {
		y := float64(y0 - float32((size+20)*i))
		vector.DrawFilledRect(screen, float32(shape.value+50), float32(y), float32(size), float32(size), shape.color, true)

		text.Draw(screen, fmt.Sprintf("%.2f", shape.value), standardFontFace, 0, int(y), color.White)

		text.Draw(screen, shape.name, standardFontFace, int(shape.value+50), int(y)+int(size)/2, color.White)
	}

}
