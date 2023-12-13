package main

import "github.com/hajimehoshi/ebiten/v2"

func getImageSize(screen *ebiten.Image) (int, int) {
	size := screen.Bounds().Size()
	w, h := size.X, size.Y
	return w, h
}
