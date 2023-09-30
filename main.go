package main

import (
	"fmt"
	"github.com/hajimehoshi/ebiten/v2"
	game2 "github.com/yngvark/ebiten-test/internal/game"
)

func main() {
	fmt.Println("Hello!")

	ebiten.SetWindowSize(640, 480)
	ebiten.SetWindowTitle("Sprite Test lol")

	game := &game2.Game{
		Title: "0.0.1",
	}

	if err := ebiten.RunGame(game); err != nil {
		panic(err)
	}
}
