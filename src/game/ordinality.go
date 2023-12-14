package main

import (
	"fmt"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text"
	"github.com/spwashi/golang-images/src/utils"
	"image/color"
)

var standardFontFace, _ = utils.LoadFont(20)
var tinyFontFace, _ = utils.LoadFont(17)

func DrawGamePosition(screen *ebiten.Image, g *Game) {
	avatarX, avatarY := getAvatarPos(g, screen)
	w, h := getImageSize(screen)

	textPosX := w * 3 / 4
	dy := 30
	yPos := 0
	face := tinyFontFace
	yPos += dy
	text.Draw(screen, fmt.Sprintf("x: %d, y: %d", int(avatarX), int(avatarY)), face, textPosX, yPos, color.White)
	yPos += dy
	text.Draw(screen, fmt.Sprintf("w: %d, h: %d", w, h), face, textPosX, yPos, color.White)
	yPos += dy
	text.Draw(screen, fmt.Sprintf("c: %d", g.counter), face, textPosX, yPos, color.White)
	yPos += dy
	text.Draw(screen, fmt.Sprintf("z: %f", g.zoom), face, textPosX, yPos, color.White)
	yPos += dy
	text.Draw(screen, fmt.Sprintf("log: %s", g.logString), face, textPosX, yPos, color.White)
}
