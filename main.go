package main

import (
	"fmt"
	"github.com/hajimehoshi/ebiten/v2"
	game2 "github.com/yngvark/ebiten-test/internal/game"
	"os"
)

func main() {
	fmt.Println("Hello!")

	var hey string
	e, got := os.LookupEnv("HEY")
	if got {
		fmt.Println("HEY:", e)
		hey = e
	} else {
		fmt.Println("not found: HEY")
		hey = "not found: HEY"
	}

	ebiten.SetWindowSize(640, 480)
	ebiten.SetWindowTitle("Sprite Test lol")

	game := &game2.Game{
		Hey: hey,
	}

	if err := ebiten.RunGame(game); err != nil {
		panic(err)
	}
}
