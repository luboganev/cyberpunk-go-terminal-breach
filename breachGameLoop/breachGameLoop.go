package breachGameLoop

import (
	"fmt"
	"log"

	"github.com/pkg/term"
)

// Exported

type OnUseBreachHole func(hoverRowIndex int, hoverColumnIndex int) bool
type DrawGameState func(hoverRowIndex int, hoverColumnIndex int, currentSelectionModeRow bool, printedLinesCount *int) bool

func RunGame(breachSurfaceSize int, drawGameState DrawGameState, onUseBreachHole OnUseBreachHole) {
	hoverRowIndex := 0
	hoverColumnIndex := 0
	printedLinesCount := 0
	currentSelectionModeRow := true

	for {
		if !drawGameState(hoverRowIndex, hoverColumnIndex, currentSelectionModeRow, &printedLinesCount) {
			break
		}
		keyCode := getInput()
		if keyCode == escape {
			return
		} else if keyCode == enter {
			if onUseBreachHole(hoverRowIndex, hoverColumnIndex) {
				currentSelectionModeRow = !currentSelectionModeRow
			}
		} else if keyCode == up && !currentSelectionModeRow {
			hoverRowIndex = (hoverRowIndex + breachSurfaceSize - 1) % breachSurfaceSize
		} else if keyCode == down && !currentSelectionModeRow {
			hoverRowIndex = (hoverRowIndex + 1) % breachSurfaceSize
		} else if keyCode == right && currentSelectionModeRow {
			hoverColumnIndex = (hoverColumnIndex + 1) % breachSurfaceSize
		} else if keyCode == left && currentSelectionModeRow {
			hoverColumnIndex = (hoverColumnIndex + breachSurfaceSize - 1) % breachSurfaceSize
		}
		// If we're gonna redraw, we need to move the cursor back up the number of lines that need redrawing
		fmt.Printf("\033[%dA", printedLinesCount)
		printedLinesCount = 0
	}
}

// Local

// UI interaction
// Raw input keycodes
var up byte = 65    // A
var down byte = 66  // B
var right byte = 67 // C
var left byte = 68  // D
var escape byte = 27
var enter byte = 13
var keys = map[byte]bool{
	up:    true,
	down:  true,
	right: true,
	left:  true,
}

// getInput will read raw input from the terminal
// It returns the raw ASCII value inputted
func getInput() byte {
	t, _ := term.Open("/dev/tty")

	err := term.RawMode(t)
	if err != nil {
		log.Fatal(err)
	}

	var read int
	readBytes := make([]byte, 3)
	read, err = t.Read(readBytes)

	t.Restore()
	t.Close()

	// Arrow keys are prefixed with the ANSI escape code which take up the first two bytes.
	// The third byte is the key specific value we are looking for.
	// For example the left arrow key is '<esc>[A' while the right is '<esc>[C'
	// See: https://en.wikipedia.org/wiki/ANSI_escape_code
	if read == 3 {
		if _, ok := keys[readBytes[2]]; ok {
			return readBytes[2]
		}
	} else {
		return readBytes[0]
	}

	return 0
}
