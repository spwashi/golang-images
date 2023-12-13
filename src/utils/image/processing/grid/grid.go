package grid

import (
	"github.com/spwashi/golang-images/src/utils/image/processing"
	"image"
	"image/color"
	"math"
)

func MakePixelGrid(img image.Image, bounds image.Rectangle, quantX int, quantY int) [][]color.RGBA64 {
	maxX := bounds.Max.X
	maxY := bounds.Max.Y

	grid := make([][]color.RGBA64, maxY)
	for y := 0; y < maxY; y++ {
		grid[y] = make([]color.RGBA64, maxX)
		for x := 0; x < maxX; x++ {
			gridX := int(math.Round(float64(x/quantX)) * float64(quantX))
			gridY := int(math.Round(float64(y/quantY)) * float64(quantY))
			pixel := img.At(gridX, gridY)
			grid[y][x] = processing.PreprocessPixel(pixel)
		}
	}

	return grid
}
