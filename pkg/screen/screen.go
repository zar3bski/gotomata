package screen

import (
	"bytes"
	"log"
	"math/rand"
	"os"
	"os/exec"
)

type Screen struct {
	State         [][]bool
	Height, Width int
}

func GetTerminalSize() []byte {
	cmd := exec.Command("stty", "size")
	cmd.Stdin = os.Stdin
	out, err := cmd.Output()
	if err != nil {
		log.Fatal(err)
	}
	return out
}

func NewScreen(width int, height int, mode string) Screen {
	// random generator
	initial_state := make([][]bool, height)
	for w := range initial_state {
		initial_state[w] = make([]bool, height)
	}
	if mode == "random" {
		for h := 0; h < height; h++ {
			for w := 0; w < height; w++ {
				initial_state[h][w] = rand.Int31()&0x01 == 0
			}
		}
	}
	return Screen{Height: height, Width: width, State: initial_state}
}

func (f *Screen) Activated(line, col int) bool {
	if line >= 0 && col >= 0 && line < f.Height && col < f.Width {
		return f.State[line][col]
	} else {
		return false
	}
}

func (s *Screen) Display() string {
	var buf bytes.Buffer
	for y := 0; y < s.Height; y++ {
		for x := 0; x < s.Width; x++ {
			b := byte(' ')
			if s.State[y][x] {
				b = '#'
			}
			buf.WriteByte(b)
		}
		buf.WriteByte('\n')
	}
	return buf.String()
}
