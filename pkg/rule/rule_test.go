package rule

import (
	"testing"

	"github.com/zar3bski/gotomata/pkg/screen"
)

func TestGameOfLife(t *testing.T) {
	// 0   0
	//   0
	//   0 0
	state := screen.ScreenState{{true, false, true}, {false, true, false}, {false, true, true}}
	new_state := GameOfLife(state)

	// expected
	//   0
	//
	//   0 0
	if new_state[0][0] != false {
		t.Fatalf(`error`)
	}
}
