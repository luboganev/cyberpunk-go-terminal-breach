package breachUI

import (
	"fmt"
	"main/breachModel"

	"github.com/jwalton/gchalk"
)

// Exported

// Prints the epic logo
func PrintLogo() {
	fmt.Println(greenOnBlackStyle("  ____      _                                 _      ____   ___ _____ _____ "))
	fmt.Println(greenOnBlackStyle(" / ___|   _| |__   ___ _ __ _ __  _   _ _ __ | | __ |___ \\ / _ \\___  |___  |"))
	fmt.Println(greenOnBlackStyle("| |  | | | | '_ \\ / _ \\ '__| '_ \\| | | | '_ \\| |/ /   __) | | | | / /   / /"))
	fmt.Println(greenOnBlackStyle("| |__| |_| | |_) |  __/ |  | |_) | |_| | | | |   <   / __/| |_| |/ /   / /"))
	fmt.Println(greenOnBlackStyle(" \\____\\__, |_.__/ \\___|_|  | .__/ \\__,_|_| |_|_|\\_\\ |_____|\\___//_/   /_/"))
	fmt.Println(greenOnBlackStyle("      |___/                  |_|"))
	fmt.Println()
}

// Prints instructions how to play
func PrintInstructions() {
	fmt.Println(greenOnBlackStyle("Use arrow keys to navigate the breach surface"))
	fmt.Println(greenOnBlackStyle("Press enter to use a breach hole"))
	fmt.Println(greenOnBlackStyle("Press escape to exit the breach protocol"))
	fmt.Println()
}

// Prints a horizontal line with the input character number length
func PrintHorizontalLine(charactersLength int, rowsCount *int) {
	for i := 0; i < charactersLength; i++ {
		fmt.Print(yellowOnBlackStyle("-"))
	}
	fmt.Println()
	*rowsCount = *rowsCount + 1
}

func PrintBreachSequenceTitle(rowsCount *int) {
	fmt.Println(yellowOnBlackStyle("SEQUENCES REQUIRED TO UPLOAD"))
	*rowsCount = *rowsCount + 1
}

// Prints the breach sequence
func PrintBreachSequence(sequence []string, breachBuffer []string, rowsCount *int) int {
	sequenceOffset, matchingAddressesCount := matchAddresses(sequence, breachBuffer)

	// if the sequence is fully matched we can count this as success
	if matchingAddressesCount == len(sequence) {
		printBreachCompleteResult(true, len(breachBuffer), rowsCount)
		return 1
	}

	// part of the sequence is matched, but we ran out of buffer before being able to match it completely
	// This is a failure
	if (len(breachBuffer) - sequenceOffset) < len(sequence) {
		printBreachCompleteResult(false, len(breachBuffer), rowsCount)
		return 1
	}

	// check if the whole buffer is already full
	isBreachBufferFull := true
	for i := 0; i < len(breachBuffer); i++ {
		if breachBuffer[i] == "--" {
			isBreachBufferFull = false
			break
		}
	}

	// if the buffer is full and we didn't match any part of the sequence, we failed
	if matchingAddressesCount == 0 && isBreachBufferFull {
		printBreachCompleteResult(false, len(breachBuffer), rowsCount)
		return 1
	}

	// start by offsetting the sequence so that it aligns with the current buffer
	for k := 0; k < sequenceOffset*3; k++ {
		fmt.Print(blackOnBlackStyle(" "))
	}
	// we then print the sequence with the matching addresses highlighted
	for k := 0; k < len(sequence); k++ {
		if k < matchingAddressesCount {
			fmt.Print(cyanOnBlackStyle(sequence[k]))
		} else {
			fmt.Print(yellowOnBlackStyle(sequence[k]))
		}
		fmt.Print(blackOnBlackStyle(" "))
	}
	fmt.Println()
	*rowsCount = *rowsCount + 1
	return 0
}

// Prints the breach buffer
func PrintBreachBuffer(breachBuffer []string, rowsCount *int) {
	fmt.Println(yellowOnBlackStyle("BUFFER"))
	*rowsCount = *rowsCount + 1
	for i := 0; i < len(breachBuffer); i++ {
		fmt.Print(yellowOnBlackStyle(breachBuffer[i]))
		fmt.Print(blackOnBlackStyle(" "))
	}
	fmt.Println()
	*rowsCount = *rowsCount + 1
}

// Prints the breach surface with the borders
func PrintBreachSurface(breachSurface [][]*breachModel.BreachHole, hoverRowIndex int, hoverColumnIndex int, currentSelectionModeRow bool, rowsCount *int) {
	var breachSurfaceSize = len(breachSurface)

	// left line + left padding + each breach hole address + each padding + right line
	PrintHorizontalLine(1+1+breachSurfaceSize*3+1, rowsCount)

	for i := 0; i < breachSurfaceSize; i++ {
		printVerticalLine()
		printEmptySpace(1)
		for j := 0; j < breachSurfaceSize; j++ {
			var isHighlighted = i == hoverRowIndex && j == hoverColumnIndex
			var isSelectable bool
			var isProjected bool
			if currentSelectionModeRow {
				isSelectable = i == hoverRowIndex
				isProjected = j == hoverColumnIndex
			} else {
				isSelectable = j == hoverColumnIndex
				isProjected = i == hoverRowIndex
			}
			printBreachHole(*breachSurface[i][j], isHighlighted, isSelectable, isProjected)
			if isSelectable && currentSelectionModeRow {
				fmt.Print(gchalk.WithBgRGB(80, 80, 80).Black(" "))
			} else if isProjected && !currentSelectionModeRow {
				fmt.Print(gchalk.WithBgRGB(50, 50, 0).Black(" "))
			} else {
				printEmptySpace(1)
			}
		}
		printVerticalLine()
		fmt.Println()
		*rowsCount = *rowsCount + 1
	}

	// left line + left padding + each breach hole address + each padding + right line
	PrintHorizontalLine(1+1+breachSurfaceSize*3+1, rowsCount)
}

// Prints the empty board space
func printEmptySpace(characterCount int) {
	for i := 0; i < characterCount; i++ {
		fmt.Print(gchalk.WithBgBlack().Black(" "))
	}
}

// Local
var greenOnBlackStyle = gchalk.WithReset().WithBgBlack().Green
var blackOnRedStyle = gchalk.WithReset().WithBgRGB(255, 87, 80).Black
var blackOnGreenStyle = gchalk.WithReset().WithBgRGB(27, 213, 118).Black
var blackOnBlackStyle = gchalk.WithReset().WithBgBlack().Black
var yellowOnBlackStyle = gchalk.WithReset().WithBgBlack().RGB(223, 240, 119)
var yellowOnGreyStyle = gchalk.WithReset().WithBgRGB(68, 75, 91).RGB(223, 240, 119)
var yellowOnYellowDimStyle = gchalk.WithReset().WithBgRGB(50, 50, 0).RGB(223, 240, 119)
var blackOnCyanStyle = gchalk.WithReset().WithBgRGB(119, 228, 226).Black
var blackOnGreyStyle = gchalk.WithReset().WithBgRGB(68, 75, 91).Black
var greyOnBlackStyle = gchalk.WithReset().WithBgBlack().RGB(68, 75, 91)
var cyanOnBlackStyle = gchalk.WithReset().WithBgBlack().RGB(119, 228, 226)

// Prints the breah hole
func printBreachHole(breachHole breachModel.BreachHole, isFocused bool, isSelectable bool, isProjected bool) {
	switch {
	case !breachHole.IsFree && isFocused:
		fmt.Print(blackOnGreyStyle("[]"))
	case !breachHole.IsFree:
		fmt.Print(greyOnBlackStyle("[]"))
	case isFocused:
		fmt.Print(blackOnCyanStyle(breachHole.Address))
	case isSelectable:
		fmt.Print(yellowOnGreyStyle(breachHole.Address))
	case isProjected:
		fmt.Print(yellowOnYellowDimStyle(breachHole.Address))
	default:
		fmt.Print(yellowOnBlackStyle(breachHole.Address))
	}
}

// Prints the vertical line
func printVerticalLine() {
	fmt.Print(yellowOnBlackStyle("|"))
}

func printBreachCompleteResult(success bool, breachBufferSize int, rowsCount *int) {
	var message string
	if success {
		message = "INSTALLED"
		fmt.Print(blackOnGreenStyle(message))
	} else {
		message = "FAILED"
		fmt.Print(blackOnRedStyle(message))
	}

	for i := 0; i < breachBufferSize*3-len(message); i++ {
		if success {
			fmt.Print(blackOnGreenStyle(" "))
		} else {
			fmt.Print(blackOnRedStyle(" "))
		}
	}
	fmt.Println()
	*rowsCount = *rowsCount + 1
}

func matchAddresses(sequence []string, breachBuffer []string) (offset int, matchesCount int) {
	offset = 0
	matchesCount = 0

breachBufferLoopLabel:
	for i := 0; i < len(breachBuffer); i++ {
	sequenceLoopLabel:
		for j := 0; j < len(sequence); j++ {
			switch {
			case i+j >= len(breachBuffer):
				break breachBufferLoopLabel
			case sequence[j] == breachBuffer[i+j]:
				matchesCount++
			case breachBuffer[i+j] == "--":
				return
			default:
				matchesCount = 0
				offset = i + 1
				break sequenceLoopLabel
			}
		}
		if matchesCount == len(sequence) {
			// We have a full match, no need to continue
			break
			// If we didn't fine any matches, there should be also no offset
		}
	}
	if matchesCount == 0 && offset != 0 {
		offset = 0
	}
	return
}
