package main

import (
	"fmt"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/audio"
	_ "image/jpeg"
	"math"
	"time"
)

type Game struct {
	active  bool
	counter int
	x, y    float64
	vx, vy  float64
	zoom    float64

	// debugging
	logString string

	// scene
	avatarImg *ebiten.Image

	// consider keyboard
	latestPausePress time.Time

	// consider touches
	touches  map[ebiten.TouchID]*touch
	touchIDs []ebiten.TouchID
	pinch    *pinch
	pan      *pan
	taps     []tap

	// audio
	audioPlayer *audio.Player
}

func NewGame() *Game {
	return &Game{
		counter:   0,
		zoom:      1,
		logString: "",
		x:         1,
		y:         1,
		vx:        5,
		vy:        0,
		active:    true,
	}
}

func (g *Game) Update() error {
	isShifted := ebiten.IsKeyPressed(ebiten.KeyShift)
	dx, dy := ebiten.Wheel()
	number := CheckNumericKeys()

	ConsiderTouches(g)

	// rate limit
	if time.Since(g.latestPausePress) < 100*time.Millisecond {
		return nil
	}

	// reset
	if ebiten.IsKeyPressed(ebiten.KeyR) {
		g.x = 1
		g.y = 1
		g.zoom = 1
		g.counter = 0
		g.logString = ""
		return nil
	}

	// play/pause
	if ebiten.IsKeyPressed(ebiten.KeySpace) {
		g.bonk()
		g.active = !g.active
		g.latestPausePress = time.Now()
	}

	// move
	delta := 1.0

	if isShifted {
		delta = 10.0

		// handle zoom
		if number != 0 {
			g.zoom = float64(number / 10)
			g.logString = fmt.Sprintf("zoom: %f", g.zoom)
			return nil
		} else if dy != 0 {
			g.zoom = math.Abs(g.zoom + (dy * 0.01))
			g.logString = fmt.Sprintf("zoom: %f", g.zoom)
			return nil
		}
	} else if ebiten.IsKeyPressed(ebiten.KeyAlt) {
		delta = .01
	}

	if ebiten.IsKeyPressed(ebiten.KeyArrowLeft) {
		g.x -= delta
	}
	if ebiten.IsKeyPressed(ebiten.KeyArrowRight) {
		g.x += delta
	}
	if ebiten.IsKeyPressed(ebiten.KeyArrowUp) {
		g.y -= delta
		g.zoom -= 0.01 * delta
	}
	if ebiten.IsKeyPressed(ebiten.KeyArrowDown) {
		g.y += delta
		g.zoom += 0.01 * delta
	}

	if number != 0 {
		g.x = number * 100
		g.logString = fmt.Sprintf("xpos: %f", g.x)
	} else {
		g.x += (dx) * 10
	}

	if g.x < 0 || g.x > 1000 {
		g.vx *= -1
	}

	if !g.active {
		return nil
	} else {
		g.x += g.vx
	}

	g.counter += 1
	return nil
}

func CheckNumericKeys() float64 {
	number := 0
	// set x from number keys
	if ebiten.IsKeyPressed(ebiten.KeyDigit0) || ebiten.IsKeyPressed(ebiten.KeyNumpad0) {
		number = 10
	}
	if ebiten.IsKeyPressed(ebiten.KeyDigit1) || ebiten.IsKeyPressed(ebiten.KeyNumpad1) {
		number = 1
	}
	if ebiten.IsKeyPressed(ebiten.KeyDigit2) || ebiten.IsKeyPressed(ebiten.KeyNumpad2) {
		number = 2
	}
	if ebiten.IsKeyPressed(ebiten.KeyDigit3) || ebiten.IsKeyPressed(ebiten.KeyNumpad3) {
		number = 3
	}
	if ebiten.IsKeyPressed(ebiten.KeyDigit4) || ebiten.IsKeyPressed(ebiten.KeyNumpad4) {
		number = 4
	}
	if ebiten.IsKeyPressed(ebiten.KeyDigit5) || ebiten.IsKeyPressed(ebiten.KeyNumpad5) {
		number = 5
	}
	if ebiten.IsKeyPressed(ebiten.KeyDigit6) || ebiten.IsKeyPressed(ebiten.KeyNumpad6) {
		number = 6
	}
	if ebiten.IsKeyPressed(ebiten.KeyDigit7) || ebiten.IsKeyPressed(ebiten.KeyNumpad7) {
		number = 7
	}
	if ebiten.IsKeyPressed(ebiten.KeyDigit8) || ebiten.IsKeyPressed(ebiten.KeyNumpad8) {
		number = 8
	}
	if ebiten.IsKeyPressed(ebiten.KeyDigit9) || ebiten.IsKeyPressed(ebiten.KeyNumpad9) {
		number = 9
	}
	return float64(number)
}

func (g *Game) bonk() {
	if g.audioPlayer == nil {
		audioPlayer := initAudioContext()
		defer audioPlayer.Close()
		g.audioPlayer = audioPlayer
	}
	//g.audioPlayer.Rewind()
	//g.audioPlayer.Play()
}

func (g *Game) Draw(screen *ebiten.Image) {
	//DrawBackground(screen, g)
	DrawGamePosition(screen, g)
	DrawAvatar(screen, g)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return outsideWidth, outsideHeight
}

func main() {
	ebiten.SetWindowSize(ScreenWidth, ScreenHeight)
	ebiten.SetWindowTitle("Move the Square")
	ebiten.SetFullscreen(true)
	ebiten.SetWindowResizingMode(ebiten.WindowResizingModeEnabled)
	game := NewGame()
	if err := ebiten.RunGame(game); err != nil {
		panic(err)
	}
}
