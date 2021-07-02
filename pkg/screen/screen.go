package screen

type Screen struct {
	State         [][]bool
	Height, Width int
}

func (f *Screen) Activated(line, col int) bool {
	if line >= 0 && col >= 0 && line < f.Height && col < f.Width {
		return f.State[line][col]
	} else {
		return false
	}
}
