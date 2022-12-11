package game

import (
	"bytes"
	_ "embed"
	"fmt"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"image"
)

//go:embed gopher.png
var Gopher_png []byte

var ebitenImage *ebiten.Image

type Game struct {
	inited bool
}

func (g Game) init() error {
	defer func() {
		g.inited = true
	}()

	img, _, err := image.Decode(bytes.NewReader(Gopher_png))
	if err != nil {
		return fmt.Errorf("decoding image: %w", err)
	}

	origEbitenImage := ebiten.NewImageFromImage(img)
	w, h := origEbitenImage.Size()

	ebitenImage = ebiten.NewImage(w, h)
	ebitenImage.DrawImage(origEbitenImage, nil)

	return nil
}

func (g Game) Update() error {
	if !g.inited {
		err := g.init()
		if err != nil {
			return fmt.Errorf("initializing: %w", err)
		}
	}

	return nil
}

func (g Game) Draw(screen *ebiten.Image) {
	ebitenutil.DebugPrint(screen, "Hello, World!")
	screen.DrawImage(ebitenImage, nil)
}

func (g Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 320, 240
}
