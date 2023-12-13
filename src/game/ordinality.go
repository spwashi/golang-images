package main

import (
	"fmt"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text"
	"github.com/spwashi/golang-images/src/utils"
	"image/color"
)

func DrawGamePosition(screen *ebiten.Image, g *Game) {
	var face, _ = utils.LoadFont(30)
	x, y := getAvatarPos(g, screen)
	text.Draw(screen, fmt.Sprintf("x: %d, y: %d", int(x), int(y)), face, 30, 30, color.White)
	w, h := getImageSize(screen)
	text.Draw(screen, fmt.Sprintf("w: %d, h: %d", w, h), face, 30, 60, color.White)
	text.Draw(screen, fmt.Sprintf("c: %d", g.counter), face, 30, 90, color.White)
	text.Draw(screen, fmt.Sprintf("z: %f", g.zoom), face, 30, 120, color.White)
	text.Draw(screen, fmt.Sprintf("l: %s", g.log), face, 30, 150, color.White)
}
