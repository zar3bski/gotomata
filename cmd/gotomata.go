package main

import (
	"fmt"

	"github.com/zar3bski/gotomata/pkg/rule"
	"github.com/zar3bski/gotomata/pkg/screen"
)

func main() {
	//TODO: work on display sizing
	screen := screen.NewScreen(200, 200, "random")
	for i := 0; i < 40; i++ {
		fmt.Print("\x0c", screen.Display())
		rule.GameOfLife(screen)
	}
}
