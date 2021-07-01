package screen

type Screen struct {
	Height int
	Width  int
}

type ScreenState [][]bool

func NewScreenState(screen Screen) ScreenState {
	screenarray := make([][]bool, screen.Height)
	for i, _ := range screenarray {
		screenarray[i] = make([]bool, screen.Width)
	}
	return screenarray
}
