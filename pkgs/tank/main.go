package main

import (
	"fmt"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
)

const (
	screenWidth       = 640
	screenHeight      = 480
	tankSpeed         = 2
	bulletSpeed       = 5
	changeDirInterval = 30 // 每 30 帧改变一次方向
	shootInterval     = 5  // 每 5 帧射击一次
)

func main() {
	ebiten.SetWindowSize(screenWidth, screenHeight)
	ebiten.SetWindowTitle("Tank Game")
	game := NewGame()
	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
	fmt.Println("Game Over")
}
