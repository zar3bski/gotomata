package main

import (
	"fmt"
	"time"

	"github.com/zar3bski/gotomata/pkg/rule"
	"github.com/zar3bski/gotomata/pkg/screen"
)

func main() {
	heigth, width := screen.GetTerminalSize()

	screen := screen.NewScreen(width, heigth-1, "random")

	for i := 0; i < 40; i++ {
		fmt.Print("\x0c", screen.Display())
		rule.GameOfLife(screen)
		time.Sleep(1 * time.Second)
	}

}
