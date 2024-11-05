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
	var result = 0
	var sequenceOffset = 0
	var matchingAddressesCount = 0

	succesMessage := "INSTALLED"
	errorMessage := "FAILED"

	var bufferAddressesCount = 0
	for j := 0; j < len(breachBuffer); j++ {
		if breachBuffer[j] != "--" {
			bufferAddressesCount++
		}
	}
	var i = 0
	for {
		if (matchingAddressesCount == len(sequence)) {
			fmt.Print(gchalk.WithBgBrightGreen().Black(succesMessage))
			for k := 0; k < (len(breachBuffer) * 3) - len(succesMessage); k++ {
				fmt.Print(gchalk.WithBgBrightGreen().Black(" "))
			}
			fmt.Println()
			*rowsCount = *rowsCount + 1
			return 1
		}
		if (i >= len(breachBuffer)) {
			break
		}
		for j := 0; j < len(sequence) && i + j < bufferAddressesCount; j++ {
			if sequence[j] == breachBuffer[i + j] {
				matchingAddressesCount++
			} else {
				matchingAddressesCount = 0
				sequenceOffset++
				break
			}
		}
		i++
	}
	
	if (len(breachBuffer) - sequenceOffset) < len(sequence) {
		fmt.Print(gchalk.WithBgBrightRed().Black(errorMessage))
		for k := 0; k < (len(breachBuffer) * 3) - len(errorMessage); k++ {
			fmt.Print(gchalk.WithBgBrightRed().Black(" "))
		}
		fmt.Println()
		*rowsCount = *rowsCount + 1
		return 1
	}

	for k := 0; k < sequenceOffset * 3; k++ {
		fmt.Print(gchalk.WithBgBlack().Yellow(" "))
	}
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
	return result
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
func PrintBreachSurface(breachSurface [][]*breachModel.BreachHole, hoverRowIndex int, hoverColumnIndex int, rowsCount *int) {
	var breachSurfaceSize = len(breachSurface)

	// left line + left padding + each breach hole address + each padding + right line
	PrintHorizontalLine(1 + 1 + breachSurfaceSize * 3 + 1, rowsCount)

	for i := 0; i < breachSurfaceSize; i++ {
		printVerticalLine()
		printEmptySpace()
		for j := 0; j < breachSurfaceSize; j++ {
			var isHighlighted = i == hoverRowIndex && j == hoverColumnIndex
			printBreachHole(*breachSurface[i][j], isHighlighted)
			printEmptySpace()
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
func printBreachHole(breachHole breachModel.BreachHole, isFocused bool) {
	switch {
	case !breachHole.IsFree && isFocused:
		fmt.Print(gchalk.WithBgBrightBlack().Black("[]"))
	case !breachHole.IsFree:
		fmt.Print(gchalk.WithBgBlack().BrightBlack("[]"))
	case isFocused:
		fmt.Print(gchalk.WithBgCyan().Black(breachHole.Address))
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

