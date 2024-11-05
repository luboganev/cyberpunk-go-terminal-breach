package breachUI

import (
	"fmt"
	"github.com/jwalton/gchalk"
	"main/breachModel"
)

// Exported

// Prints the epic logo
func PrintLogo() {
	fmt.Println(gchalk.WithBgBlack().Green(" .o88b. db    db d8888b. d88888b d8888b. d8888b. db    db d8b   db db   dD     .d888b.  .d88b.  d88888D d88888D "))
	fmt.Println(gchalk.WithBgBlack().Green("d8P  Y8 `8b  d8' 88  `8D 88'     88  `8D 88  `8D 88    88 888o  88 88 ,8P'     VP  `8D .8P  88. VP  d8' VP  d8'"))
	fmt.Println(gchalk.WithBgBlack().Green("8P       `8bd8'  88oooY' 88ooooo 88oobY' 88oodD' 88    88 88V8o 88 88,8P          odD' 88  d'88    d8'     d8' "))
	fmt.Println(gchalk.WithBgBlack().Green("8b         88    88~~~b. 88~~~~~ 88`8b   88~~~   88    88 88 V8o88 88`8b        .88'   88 d' 88   d8'     d8'  "))
	fmt.Println(gchalk.WithBgBlack().Green("Y8b  d8    88    88   8D 88.     88 `88. 88      88b  d88 88  V888 88 `88.     j88.    `88  d8'  d8'     d8'   "))
	fmt.Println(gchalk.WithBgBlack().Green(" `Y88P'    YP    Y8888P' Y88888P 88   YD 88      ~Y8888P' VP   V8P YP   YD     888888D  `Y88P'  d8'     d8'    "))
	fmt.Println()
	fmt.Println()
}

// Prints instructions how to play
func PrintInstructions() {
	fmt.Println(gchalk.WithBgBlack().Green("Use arrow keys to navigate the breach surface"))
	fmt.Println(gchalk.WithBgBlack().Green("Press enter to use a breach hole"))
	fmt.Println(gchalk.WithBgBlack().Green("Press escape to exit the breach protocol"))
	fmt.Println()
	fmt.Println()
}

// Prints a horizontal line with the input character number length
func PrintHorizontalLine(charactersLength int, rowsCount *int) {
	for i := 0; i < charactersLength; i++ {
		fmt.Print(gchalk.WithBgBlack().Yellow("-"))
	}
	fmt.Println()
	*rowsCount = *rowsCount + 1
}

func PrintBreachSequenceTitle(rowsCount *int) {
	fmt.Println(gchalk.WithBgBlack().Yellow("SEQUENCES REQUIRED TO UPLOAD"))
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
	for k := 0; k < sequenceOffset * 3; k++ {
		fmt.Print(gchalk.WithBgBlack().Yellow(" "))
	}
	// we then print the sequence with the matching addresses highlighted
	for k := 0; k < len(sequence); k++ {
		if (k < matchingAddressesCount) {
			fmt.Print(gchalk.WithBgBlack().Cyan(sequence[k]))
		} else {
			fmt.Print(gchalk.WithBgBlack().Yellow(sequence[k]))
		}
		fmt.Print(gchalk.WithBgBlack().Yellow(" "))
	}
	fmt.Println()
	*rowsCount = *rowsCount + 1
	return 0
}

// Prints the breach buffer
func PrintBreachBuffer(breachBuffer []string, rowsCount *int) {
	fmt.Println(gchalk.WithBgBlack().Yellow("BUFFER"))
	*rowsCount = *rowsCount + 1
	for i := 0; i < len(breachBuffer); i++ {
		fmt.Print(gchalk.WithBgBlack().Yellow(breachBuffer[i]))
		fmt.Print(gchalk.WithBgBlack().Yellow(" "))
	}
	fmt.Println()
	*rowsCount = *rowsCount + 1
}

// Prints the breach surface with the borders
func PrintBreachSurface(breachSurface [][]*breachModel.BreachHole, hoverRowIndex int, hoverColumnIndex int, currentSelectionModeRow bool, rowsCount *int) {
	var breachSurfaceSize = len(breachSurface)

	// left line + left padding + each breach hole address + each padding + right line
	PrintHorizontalLine(1 + 1 + breachSurfaceSize * 3 + 1, rowsCount)

	for i := 0; i < breachSurfaceSize; i++ {
		printVerticalLine()
		printEmptySpace()
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
				printEmptySpace()
			}
		}
		printVerticalLine()
		fmt.Println()
		*rowsCount = *rowsCount + 1
	}

	// left line + left padding + each breach hole address + each padding + right line
	PrintHorizontalLine(1 + 1 + breachSurfaceSize * 3 + 1, rowsCount)
}

// Local

// Prints the breah hole
func printBreachHole(breachHole breachModel.BreachHole, isFocused bool, isSelectable bool, isProjected bool) {
	switch {
	case !breachHole.IsFree && isFocused:
		fmt.Print(gchalk.WithBgBrightBlack().Black("[]"))
	case !breachHole.IsFree:
		fmt.Print(gchalk.WithBgBlack().BrightBlack("[]"))
	case isFocused:
		fmt.Print(gchalk.WithBgCyan().Black(breachHole.Address))
	case isSelectable:
		fmt.Print(gchalk.WithBgRGB(80, 80, 80).Yellow(breachHole.Address))
	case isProjected:
		fmt.Print(gchalk.WithBgRGB(50, 50, 0).Yellow(breachHole.Address))
	default:
		fmt.Print(gchalk.WithBgBlack().Yellow(breachHole.Address))
	}
}

// Prints the vertical line
func printVerticalLine() {
	fmt.Print(gchalk.WithBgBlack().Yellow("|"))
}

// Prints the empty board space
func printEmptySpace() {
	fmt.Print(gchalk.WithBgBlack().Black(" "))
}

func printBreachCompleteResult(success bool, breachBufferSize int, rowsCount *int) {
	var message string 
	if success {
		message = "INSTALLED"
	} else {
		message = "FAILED"
	}
	if success {
		fmt.Print(gchalk.WithBgBrightGreen().Black(message))
	} else {
		fmt.Print(gchalk.WithBgBrightRed().Black(message))
	}
	for i := 0; i < breachBufferSize * 3 - len(message); i++ {
		if success {
			fmt.Print(gchalk.WithBgBrightGreen().Black(" "))
		} else {
			fmt.Print(gchalk.WithBgBrightRed().Black(" "))
		}	
	}
	fmt.Println()
	*rowsCount = *rowsCount + 1
}

func matchAddresses(sequence []string, breachBuffer []string) (int, int) {
	var offset = 0
	var matchesCount = 0

	breachBufferLoopLabel:
	for i := 0; i < len(breachBuffer); i++ {
		sequenceLoopLabel:
		for j := 0; j < len(sequence); j++ {
			switch {
			case i + j >= len(breachBuffer):
				break breachBufferLoopLabel
			case sequence[j] == breachBuffer[i + j]:
				matchesCount++
			case breachBuffer[i + j] == "--":
				return offset, matchesCount
			default:
				matchesCount = 0
				offset = i + 1
				break sequenceLoopLabel
			}
		}
		if (matchesCount == len(sequence)) {
			// We have a full match, no need to continue
			break
		}
	}
	// If we didn't fine any matches, there should be also no offset
	if (matchesCount == 0 && offset != 0) {
		offset = 0
	}
	return offset, matchesCount
}

