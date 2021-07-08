package main

import (
	"fmt"
	"time"

	"github.com/zar3bski/gotomata/pkg/rule"
	"github.com/zar3bski/gotomata/pkg/screen"
)

func main() {
	//TODO: work on display sizing https://stackoverflow.com/questions/16569433/get-terminal-size-in-go
	heigth, width := screen.GetTerminalSize()

	screen := screen.NewScreen(width, heigth, "random")
	fmt.Printf("Height: %d\n", heigth)
	fmt.Printf("Width: %d\n", width)

	for i := 0; i < 2; i++ {
		fmt.Print("\x0c", screen.Display())
		rule.GameOfLife(screen)
		time.Sleep(1 * time.Second)
	}
	// TODO: identify height and width
	//fmt.Printf("out: %v\n", dimensions)
}
