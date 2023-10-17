package processing

import "image/color"

func PreprocessPixel(pixel color.Color) color.RGBA64 {
	r, g, b, _ := pixel.RGBA()

	newR := uint16(r)
	newG := uint16(g)
	newB := uint16(b)
	newA := uint16(255)

	if uint8(newR) > 255 {
		tempA := newR
		tempB := newG
		tempC := newB
		newR = tempA
		newG = tempB
		newB = tempC
	}

	return color.RGBA64{
		R: newR,
		G: newG,
		B: newB,
		A: newA,
	}
}

func PostProcessPixel(pixel color.RGBA64) color.RGBA64 {
	return pixel
}
