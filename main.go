package main

import (
	"fmt"
	"github.com/hajimehoshi/ebiten/v2"
	game2 "github.com/yngvark/ebiten-test/internal/game"
)

func main() {
	fmt.Println("Hello!")

	game := &game2.Game{
		Title: "0.0.1",
	}

	game.Init()

	if err := ebiten.RunGame(game); err != nil {
		panic(err)
	}
}
