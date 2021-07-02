package screen

import "testing"

func TestActivated(t *testing.T) {
	s := Screen{Height: 3, Width: 3, State: [][]bool{{true, true, true}, {false, true, false}, {false, true, true}}}
	status := s.Activated(0, 0)
	if status != true {
		t.Fatalf(`error`)
	}
	status2 := s.Activated(-1, 0)
	if status2 != false {
		t.Fatalf(`error`)
	}
	status3 := s.Activated(1, 0)
	if status3 != false {
		t.Fatalf(`error`)
	}
}
