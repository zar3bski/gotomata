package rule

import (
	"fmt"
	"testing"

	"github.com/zar3bski/gotomata/pkg/screen"
)

func TestGameOfLife(t *testing.T) {
	// 0   0
	//   0
	//   0 0
	s := screen.Screen{Height: 3, Width: 3, State: [][]bool{{true, false, true}, {false, true, false}, {false, true, true}}}

	GameOfLife(s)
	fmt.Printf("%v", s.State)
	// expected
	//   0
	//
	//   0 0
	if s.State[0][0] != false {
		t.Fatalf(`error`)
	}
	if s.State[0][1] != true {
		t.Fatalf(`error`)
	}

}
