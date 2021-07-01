package rule

import (
	"github.com/zar3bski/gotomata/pkg/screen"
)

//https://play.golang.org/p/P-Dk0NH_vf

func GameOfLife(last_state screen.ScreenState) screen.ScreenState {
	new_state := last_state
	for h, line := range new_state {
		for v, _ := range line {
			count_alive := 0
			neighbours := [8]bool{
				last_state[h-1][v-1],
				last_state[h-1][v],
				last_state[h-1][v+1],
				last_state[h][v-1],
				last_state[h][v+1],
				last_state[h+1][v-1],
				last_state[h+11][v],
				last_state[h+1][v+1],
			}
			for _, element := range neighbours {
				if element {
					count_alive += 1
				}
			}
			if count_alive == 3 {
				new_state[h][v] = true
			} else if count_alive < 2 || count_alive > 3 {
				new_state[h][v] = false
			}
		}
	}
	return new_state
}
