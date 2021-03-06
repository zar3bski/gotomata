package rule

import (
	"github.com/zar3bski/gotomata/pkg/screen"
)

func GameOfLife(screen screen.Screen) {
	new_state := screen.State
	for l, line := range new_state {
		for c, _ := range line {
			count_alive := 0
			for i := -1; i <= 1; i++ {
				for j := -1; j <= 1; j++ {
					if (j != 0 || i != 0) && screen.Activated(l+i, c+j) {
						count_alive++
					}
				}
			}
			if count_alive == 3 {
				new_state[l][c] = true
			} else if count_alive == 2 {
				new_state[l][c] = screen.Activated(l, c)
			} else {
				new_state[l][c] = false
			}
		}
	}
	copy(screen.State, new_state)
}
