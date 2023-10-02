package game

import (
	"bytes"
	_ "embed"
	"fmt"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/yngvark/ebiten-test/internal/game/tiles"
	"image"
	"image/color"
	"math/rand"
)

const tileSize = 48
const scale = 5

//go:embed gopher.png
var gopherPng []byte

// Images
var gopherImage *ebiten.Image
var middleRectangle *ebiten.Image

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

	// Gopher
	gopherImageSource, _, err := image.Decode(bytes.NewReader(gopherPng))
	if err != nil {
		return fmt.Errorf("decoding image: %w", err)
	}

	gopherImage = ebiten.NewImageFromImage(gopherImageSource)

	// Middle rectangle
	middleRectangle = ebiten.NewImage(50, 50)
	middleRectangle.Fill(color.NRGBA{R: 0x80, G: 0, B: 0, A: 0xff})

	tiles.Init()

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
	tiles.Draw(screen)
	g.drawMovingRectangle(screen)
}

func (g *Game) drawStillImage(screen *ebiten.Image) {
	screen.DrawImage(gopherImage, nil)
}

func (g *Game) drawMovingRectangle(screen *ebiten.Image) {
	screenWidth, screenHeight := screen.Size()
	rectangleWidth, rectangleHeight := middleRectangle.Size()

	// Calculate the x and y coordinates to draw the image at the center of the window.
	x := float64(screenWidth/2-rectangleWidth/2) + g.rectangleX
	y := float64(screenHeight/2-rectangleHeight/2) + g.rectangleY

	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(x, y)

	screen.DrawImage(middleRectangle, op)
}

//func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
//	return 320, 240
//}

// Layout takes the outside size (e.g., the window size) and returns the (logical) screen size.
func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return tiles.WorldMapWidth() * tileSize, tiles.WorldMapHeight() * tileSize
}

func (g *Game) Init() {
	screenWidth := tiles.WorldMapWidth() * tileSize
	screenHeight := tiles.WorldMapHeight() * tileSize

	ebiten.SetWindowSize(screenWidth*scale, screenHeight*scale)
	//ebiten.SetWindowSize(len(worldMap[0])*tileSize, len(worldMap)*tileSize)
	ebiten.SetWindowTitle("Sprite Test lol")
}
