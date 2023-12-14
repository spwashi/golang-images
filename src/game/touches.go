package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"math"
)

type touch struct {
	originX, originY int
	currX, currY     int
	duration         int
	wasPinch, isPan  bool
}

type pinch struct {
	id1, id2 ebiten.TouchID
	originH  float64
	prevH    float64
}

type pan struct {
	id ebiten.TouchID

	prevX, prevY     int
	originX, originY int
}

type tap struct {
	X, Y int
}

// distance between points a and b.
func distance(xa, ya, xb, yb int) float64 {
	x := math.Abs(float64(xa - xb))
	y := math.Abs(float64(ya - yb))
	return math.Sqrt(x*x + y*y)
}

func ConsiderTouches(g *Game) {
	// Clear the previous frame's taps.
	g.taps = g.taps[:0]

	// What touches have just ended?
	for id, t := range g.touches {
		if inpututil.IsTouchJustReleased(id) {
			g.logString = "released a touch"
			if g.pinch != nil && (id == g.pinch.id1 || id == g.pinch.id2) {
				g.pinch = nil
			}
			if g.pan != nil && id == g.pan.id {
				g.pan = nil
			}

			// If this one has not been touched long (30 frames can be assumed
			// to be 500ms), or moved far, then it's a tap.
			diff := distance(t.originX, t.originY, t.currX, t.currY)
			if !t.wasPinch && !t.isPan && (t.duration <= 30 || diff < 2) {
				g.taps = append(g.taps, tap{
					X: t.currX,
					Y: t.currY,
				})
			}

			delete(g.touches, id)
		}
	}

	// What touches are new in this frame?
	g.touchIDs = inpututil.AppendJustPressedTouchIDs(g.touchIDs[:0])
	for _, id := range g.touchIDs {
		g.logString = "at least one touch"
		x, y := ebiten.TouchPosition(id)
		g.touches[id] = &touch{
			originX: x, originY: y,
			currX: x, currY: y,
		}
	}

	g.touchIDs = ebiten.AppendTouchIDs(g.touchIDs[:0])

	// Update the current position and durations of any touches that have
	// neither begun nor ended in this frame.
	for _, id := range g.touchIDs {
		t := g.touches[id]
		t.duration = inpututil.TouchPressDuration(id)
		t.currX, t.currY = ebiten.TouchPosition(id)
	}

	// Interpret the raw touch data that's been collected into g.touches into
	// gestures like two-finger pinch or single-finger pan.
	switch len(g.touches) {
	case 2:
		// Potentially the user is making a pinch gesture with two fingers.
		// If the diff between their origins is different to the diff between
		// their currents and if these two are not already a pinch, then this is
		// a new pinch!
		g.logString = "pinch"
		id1, id2 := g.touchIDs[0], g.touchIDs[1]
		t1, t2 := g.touches[id1], g.touches[id2]
		originDiff := distance(t1.originX, t1.originY, t2.originX, t2.originY)
		currDiff := distance(t1.currX, t1.currY, t2.currX, t2.currY)
		if g.pinch == nil && g.pan == nil && math.Abs(originDiff-currDiff) > 3 {
			t1.wasPinch = true
			t2.wasPinch = true
			g.pinch = &pinch{
				id1:     id1,
				id2:     id2,
				originH: originDiff,
				prevH:   originDiff,
			}
		}
	case 1:
		// Potentially this is a new pan.
		g.logString = "pan"
		id := g.touchIDs[0]
		t := g.touches[id]
		if !t.wasPinch && g.pan == nil && g.pinch == nil {
			diff := math.Abs(distance(t.originX, t.originY, t.currX, t.currY))
			if diff > 1 {
				t.isPan = true
				g.pan = &pan{
					id:      id,
					originX: t.originX,
					originY: t.originY,
					prevX:   t.originX,
					prevY:   t.originY,
				}
			}
		}
	}

	// Copy any active pinch gesture's movement to the Game's zoom.
	if g.pinch != nil {
		x1, y1 := ebiten.TouchPosition(g.pinch.id1)
		x2, y2 := ebiten.TouchPosition(g.pinch.id2)
		curr := distance(x1, y1, x2, y2)
		delta := curr - g.pinch.prevH
		g.pinch.prevH = curr

		g.zoom += (delta / 100) * g.zoom
		if g.zoom < 0.25 {
			g.zoom = 0.25
		} else if g.zoom > 10 {
			g.zoom = 10
		}
	}

	// Copy any active pan gesture's movement to the Game's value and y pan values.
	if g.pan != nil {
		currX, currY := ebiten.TouchPosition(g.pan.id)
		deltaX, deltaY := currX-g.pan.prevX, currY-g.pan.prevY

		g.pan.prevX, g.pan.prevY = currX, currY

		g.x += float64(deltaX)
		g.y += float64(deltaY)
	}

	// If the user has tapped, then reset the Game's pan and zoom.
	if len(g.taps) > 0 {
		g.zoom = 1.0
	}
}
