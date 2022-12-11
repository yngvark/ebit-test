package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	game2 "yngvark.com/ebiten-test/internal/game"
)

func main() {
	ebiten.SetWindowSize(640, 480)
	ebiten.SetWindowTitle("Sprite Test lol")

	game := &game2.Game{}

	if err := ebiten.RunGame(game); err != nil {
		panic(err)
	}
}
