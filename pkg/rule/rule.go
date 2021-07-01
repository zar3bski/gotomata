package rule

import (
	"fmt"

	"github.com/zar3bski/gotomata/pkg/screen"
)

func GameOfLife(last_state screen.ScreenState) screen.ScreenState {
	new_state := last_state
	for _, h := range new_state {
		for _, cell := range h {
			// implement logic here
			// neighbours =
			fmt.Print(cell, " ")
		}
		fmt.Println()
	}
	return new_state
}
