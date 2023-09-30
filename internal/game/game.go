package game

import (
	"bytes"
	_ "embed"
	"fmt"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"image"
	"image/color"
	"math/rand"
)

//go:embed gopher.png
var gopherPng []byte

var ebitenImage *ebiten.Image

type Game struct {
	inited bool

	x float64
	y float64
}

func (g *Game) init() error {
	defer func() {
		g.inited = true
	}()

	img, _, err := image.Decode(bytes.NewReader(gopherPng))
	if err != nil {
		return fmt.Errorf("decoding image: %w", err)
	}

	origEbitenImage := ebiten.NewImageFromImage(img)
	w, h := origEbitenImage.Size()

	ebitenImage = ebiten.NewImage(w, h)
	ebitenImage.DrawImage(origEbitenImage, nil)

	return nil
}

func (g *Game) Update() error {
	if !g.inited {
		err := g.init()
		if err != nil {
			return fmt.Errorf("initializing: %w", err)
		}
	}

	g.x += float64(rand.Intn(3)) - 1 // This will add -1, 0, or 1 to g.x
	g.y += float64(rand.Intn(3)) - 1 // This will add -1, 0, or 1 to g.y

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	// Fill the screen with a white color.
	screen.Fill(color.White)

	// Print some text
	ebitenutil.DebugPrint(screen, "Hello, World!")

	// Draw a still image
	screen.DrawImage(ebitenImage, nil)

	// Draw a moving rectangle (x and y coordinates, width, height).
	rect := ebiten.NewImage(50, 50)
	rect.Fill(color.NRGBA{R: 0x80, G: 0, B: 0, A: 0xff})

	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(g.x, g.y)
	screen.DrawImage(rect, op)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 320, 240
}
