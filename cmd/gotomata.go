package main

import (
	"fmt"
	"time"

	"github.com/zar3bski/gotomata/pkg/rule"
	"github.com/zar3bski/gotomata/pkg/screen"
)

func main() {
	//TODO: work on display sizing https://stackoverflow.com/questions/16569433/get-terminal-size-in-go
	dimensions := screen.GetTerminalSize()
	screen := screen.NewScreen(int(dimensions[0]), int(dimensions[1]), "random")
	for i := 0; i < 2; i++ {
		fmt.Print("\x0c", screen.Display())
		rule.GameOfLife(screen)
		time.Sleep(1 * time.Second)
	}
	// TODO: identify height and width
	fmt.Printf("out: %d\n", dimensions[0])
}
