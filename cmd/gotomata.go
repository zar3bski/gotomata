package main

import (
	"fmt"

	"github.com/zar3bski/gotomata/pkg/screen"
)

func main() {

	var currentscreen = screen.Screen{Height: 8, Width: 6}

	state := screen.NewScreenState(currentscreen)
	fmt.Printf("%v", state)
}
