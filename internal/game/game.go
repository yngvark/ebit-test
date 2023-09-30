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

var gopherImage *ebiten.Image

type Game struct {
	Title string

	inited bool

	rectangleX float64
	rectangleY float64
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

	gopherImage = ebiten.NewImage(w, h)

	return nil
}

// Update proceeds the game state.
// Update is called every tick (1/60 [s] by default).
// https://pkg.go.dev/github.com/hajimehoshi/ebiten/v2#section-readme
func (g *Game) Update() error {
	if !g.inited {
		err := g.init()
		if err != nil {
			return fmt.Errorf("initializing: %w", err)
		}
	}

	g.rectangleX += float64(rand.Intn(3)) - 1
	g.rectangleY += float64(rand.Intn(3)) - 1

	return nil
}

// Draw draws the game screen.
// Draw is called every frame (typically 1/60[s] for 60Hz display).
func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(color.White)

	ebitenutil.DebugPrint(screen, g.Title)

	g.drawStillImage(screen)
	g.drawMovingRectangle(screen)
}

func (g *Game) drawStillImage(screen *ebiten.Image) {
	screen.DrawImage(gopherImage, nil)
}

func (g *Game) drawMovingRectangle(screen *ebiten.Image) {
	rect := ebiten.NewImage(50, 50)
	rect.Fill(color.NRGBA{R: 0x80, G: 0, B: 0, A: 0xff})

	screenWidth, screenHeight := screen.Size()
	rectangleWidth, rectangleHeight := rect.Size()

	// Calculate the x and y coordinates to draw the image at the center of the window.
	x := float64(screenWidth/2-rectangleWidth/2) + g.rectangleX
	y := float64(screenHeight/2-rectangleHeight/2) + g.rectangleY

	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(x, y)

	screen.DrawImage(rect, op)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 320, 240
}
