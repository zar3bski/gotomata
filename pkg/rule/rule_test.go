package rule

import (
	"reflect"
	"testing"

	"github.com/zar3bski/gotomata/pkg/screen"
)

func TestGameOfLife(t *testing.T) {
	// 0   0
	//   0
	//   0 0
	s := screen.Screen{Height: 3, Width: 3, State: [][]bool{{true, false, true}, {false, true, false}, {false, true, true}}}

	GameOfLife(s)
	expected := [][]bool{{false, true, false}, {false, false, false}, {false, true, true}}
	// expected
	//   0
	//
	//   0 0
	if reflect.DeepEqual(expected, s.State) {
		t.Fatalf(`expected %v\nbut got  %v`, expected, s.State)
	}
}
