package main

import (
	"fmt"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text"
	"github.com/hajimehoshi/ebiten/v2/vector"
	"github.com/spwashi/golang-images/src/utils"
	"golang.org/x/image/font"
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

var highlightedFontFace, _ = utils.LoadFont(50)
var bigOnameFontFace, _ = utils.LoadFont(50)
var numberFontSize = float64(50)
var numberFontFace, _ = utils.LoadFont(numberFontSize)

func drawSeekers(screen *ebiten.Image, counter float64, scale float64) {
	factor := 5
	k := math.Abs(scale / float64(factor))
	w, h := getImageSize(screen)
	h -= 50
	size := 100
	xVal := counter * k
	nColor := color.RGBA{90, 150, 90, 255}
	y0 := float32(h - size)
	shapes := []Marker{
		{name: "n!", value: math.Gamma(xVal + 1), color: color.RGBA{R: 80, G: 0, B: 80, A: 1}},
		{name: "2^n", value: math.Pow(2, xVal), color: color.RGBA{R: 0, G: 0, B: 80, A: 1}},
		{name: "n^2", value: math.Pow(xVal, 2), color: color.RGBA{R: 80, G: 0, B: 0, A: 1}},
		{name: "n log(n)", value: xVal * math.Log(xVal), color: color.RGBA{R: 0, G: 80, B: 80, A: 1}},
		{name: "n", value: math.Pow(xVal, 1), color: color.RGBA{R: 80, G: 80, B: 0, A: 1}},
		{name: "log(n)", value: math.Log(xVal), color: color.RGBA{R: 0, G: 80, B: 0, A: 1}},
	}
	rectOffset := font.MeasureString(bigOnameFontFace, "n log(n)").Ceil() + 20

	for i, shape := range shapes {
		y := float64(y0 - float32((size+int(numberFontSize))*i))
		rectWidth := float32(shape.value)
		strWidth := font.MeasureString(bigOnameFontFace, shape.name).Ceil()
		strOffset := (rectOffset - strWidth) / 2

		vector.StrokeLine(screen, 0, float32(y), float32(w), float32(y), 2, color.RGBA{R: 120, G: 120, B: 120, A: 1}, true)

		// shape name
		text.Draw(screen, shape.name, bigOnameFontFace, strOffset, int(y)+int(size)/2, color.RGBA{200, 200, 200, 200})

		// shape value
		text.Draw(screen, fmt.Sprintf("%.2f", shape.value), numberFontFace, rectOffset, int(y+numberFontSize), nColor)

		// value rect
		vector.DrawFilledRect(screen, float32(rectOffset), float32(y), rectWidth, float32(size), shape.color, true)
	}

	xStr := fmt.Sprintf("x:%d", int(counter))
	xStrWidth := font.MeasureString(highlightedFontFace, xStr).Ceil()
	kStr := fmt.Sprintf("k:%.2f", k)
	kStrWidth := font.MeasureString(highlightedFontFace, kStr).Ceil()
	nStr := fmt.Sprintf("n:%.1f", xVal)
	nStrWidth := font.MeasureString(highlightedFontFace, nStr).Ceil()

	offsetX := 10
	offsetY := 60
	gray := color.RGBA{90, 90, 100, 255}
	text.Draw(screen, xStr, highlightedFontFace, offsetX, offsetY, gray)
	offsetX = offsetX + xStrWidth + 50
	text.Draw(screen, kStr, highlightedFontFace, offsetX, offsetY, gray)
	offsetX = offsetX + kStrWidth + 50
	text.Draw(screen, nStr, highlightedFontFace, offsetX, offsetY, nColor)
	offsetX = offsetX + nStrWidth + 100
	// draw the formula k * x = n
	text.Draw(screen, "k * x =  ", highlightedFontFace, offsetX, 60, color.RGBA{90, 60, 100, 255})
	text.Draw(screen, "k   x    ", highlightedFontFace, offsetX, 60, gray)
	text.Draw(screen, "        n", highlightedFontFace, offsetX, 60, nColor)
}
