package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"log"
)

func loadImage(path string) (*ebiten.Image, error) {
	img, _, err := ebitenutil.NewImageFromFile(path)
	if err != nil {
		return nil, err
	}
	return ebiten.NewImageFromImage(img), nil

}

func DrawBackground(screen *ebiten.Image, g *Game) {
	if g.avatarImg == nil {
		log.Default().Println("loading image")
		var err error
		g.avatarImg, err = loadImage("input/image.jpg")
		if err != nil {
			panic(err)
		}
	}
	op := &ebiten.DrawImageOptions{}
	//op.GeoM.Translate(g.value, g.y)
	//op.GeoM.Scale(0.5, 0.5)
	//screenColor := color.RGBA{R: 0x0f, G: 0x9f, B: 0x9f, A: 0x5f}
	if !g.active {
		screen.DrawImage(g.avatarImg, op)
		//w, h := getImageSize(screen)
		//vector.DrawFilledRect(screen, 0, 0, float32(w/100), float32(h/100), screenColor, true)
	}
}
