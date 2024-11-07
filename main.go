package main

import (
	"fmt"
	"main/breachGameLoop"
	"main/breachModel"
	"main/breachUI"
)

func main() {
	breachSurfaceSize := 6
	breachSurface := breachModel.GenerateBreachSurface(breachSurfaceSize)
	breachSequences := breachModel.GenerateBreachSequencesFromSurface(6, breachSurface, 3)
	breachSequence1 := breachSequences[0]
	breachSequence2 := breachSequences[1]
	breachSequence3 := breachSequences[2]

	breachBuffer := breachModel.GenerateBreachBuffer(6)
	currentBufferIndex := 0

	breachUI.PrintLogo()
	breachUI.PrintInstructions()

	onUseBreachHole := func(hoverRowIndex int, hoverColumnIndex int) bool {
		currentBreachedHole := breachSurface[hoverRowIndex][hoverColumnIndex]
		if currentBreachedHole.IsFree {
			currentBreachedHole.IsFree = false
			breachBuffer[currentBufferIndex] = currentBreachedHole.Address
			currentBufferIndex++
			return true
		}
		return false
	}

	drawGameState := func(hoverRowIndex int, hoverColumnIndex int, currentSelectionModeRow bool, printedLinesCount *int) bool {
		breachBufferWidth := breachUI.PrintBreachBuffer(breachBuffer, printedLinesCount)
		fmt.Println()
		*printedLinesCount++
		breachUI.PrintBreachSequenceTitle(printedLinesCount)
		breach1Result := breachUI.PrintBreachSequence(breachSequence1, breachBuffer, breachBufferWidth, printedLinesCount)
		breach2Result := breachUI.PrintBreachSequence(breachSequence2, breachBuffer, breachBufferWidth, printedLinesCount)
		breach3Result := breachUI.PrintBreachSequence(breachSequence3, breachBuffer, breachBufferWidth, printedLinesCount)
		fmt.Println()
		*printedLinesCount++
		breachUI.PrintBreachSurface(breachSurface, hoverRowIndex, hoverColumnIndex, currentSelectionModeRow, printedLinesCount)

		if breach1Result != 0 && breach2Result != 0 && breach3Result != 0 {
			return false
		}
		if currentBufferIndex >= len(breachBuffer) {
			return false
		}
		return true
	}

	breachGameLoop.RunGame(breachSurfaceSize, drawGameState, onUseBreachHole)
}
