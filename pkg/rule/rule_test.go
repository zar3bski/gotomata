package rule

import (
	"reflect"
	"testing"

	"github.com/zar3bski/gotomata/pkg/screen"
)

func TestGameOfLife(t *testing.T) {

	s := screen.Screen{Height: 4, Width: 5, State: [][]bool{{true, true, false, true, false}, {false, true, true, false, false}, {false, false, true, false, false}, {true, true, false, false, true}}}

	GameOfLife(s)
	expected := [][]bool{{true, true, false, false, false}, {true, false, false, true, false}, {true, false, true, false, false}, {true, false, false, false, false}}

	if reflect.DeepEqual(expected, s.State) {
		t.Fatalf(`expected %v\nbut got  %v`, expected, s.State)
	}
}
