package breachModel

import (
	"math/rand"
	"time"
)

// The addresses of the breach holes that are known to hackers
var knownBreachHoleAddresses = [...]string{"1C", "55", "7A", "BD", "E9", "FF"}

const BreachBufferFreeSlotSymbol = "[]"

// Represents a breach hole with a known address that can be used to hack into a system
type BreachHole struct {
	Address string
	IsFree  bool
}

// Generates a square
func GenerateBreachSurface(size int) [][]*BreachHole {
	rand.Seed(time.Now().UnixNano())
	var breachSurface = make([][]*BreachHole, size)
	for i := 0; i < size; i++ {
		breachSurface[i] = make([]*BreachHole, size)
		for j := 0; j < size; j++ {
			randItemIndex := rand.Intn(len(knownBreachHoleAddresses))
			breachHole := BreachHole{Address: knownBreachHoleAddresses[randItemIndex], IsFree: true}
			breachSurface[i][j] = &breachHole
		}
	}
	return breachSurface
}

type BreachHoleWithPosition struct {
	hole      *BreachHole
	PositionX int
	PositionY int
}

type BreachSquenceWithNextPosition struct {
	sequence  []string
	positionX int
	positionY int
	isRow     bool
}

// Generates an array of random breach hole addresses with a length based on a known breach surface
// starting from a specific position (positionX, positionY) and in a specific direction (isRow)
func GenerateBreachSingleSequenceFromSurface(length int, surface [][]*BreachHole, positionX int, positionY int, isRow bool) BreachSquenceWithNextPosition {
	rand.Seed(time.Now().UnixNano())

	var breachSequence = make([]string, 0)
	for i := 0; i < length; i++ {
		var rowOfAvailableHoles = make([]BreachHoleWithPosition, 0)

		if isRow {
			for j := 0; j < len(surface); j++ {
				var hole = surface[positionY][j]
				if hole.IsFree {
					rowOfAvailableHoles = append(rowOfAvailableHoles, BreachHoleWithPosition{hole: hole, PositionX: j, PositionY: positionY})
				}
			}
		} else {
			for j := 0; j < len(surface); j++ {
				var hole = surface[j][positionX]
				if hole.IsFree {
					rowOfAvailableHoles = append(rowOfAvailableHoles, BreachHoleWithPosition{hole: hole, PositionX: positionX, PositionY: j})
				}
			}
		}

		var nextHole = rowOfAvailableHoles[rand.Intn(len(rowOfAvailableHoles))]
		breachSequence = append(breachSequence, nextHole.hole.Address)
		// Ensure we don't use the same hole twice
		nextHole.hole.IsFree = false
		// Update the position and direction for next iteration
		positionX = nextHole.PositionX
		positionY = nextHole.PositionY
		isRow = !isRow
	}

	return BreachSquenceWithNextPosition{
		sequence:  breachSequence,
		positionX: positionX,
		positionY: positionY,
		isRow:     isRow,
	}
}

// Generates a list of sequences with a combined total length based on a known breach surface
// Will randomly add overlap or remove parts of the sequences to make them more interesting
// Shuffles the sequences at the end to not indicate the order of the sequences
func GenerateBreachSequencesFromSurface(length int, surface [][]*BreachHole, count int) [][]string {

	var shallowCopyOfSurface = make([][]*BreachHole, len(surface))
	for i := range surface {
		shallowCopyOfSurface[i] = make([]*BreachHole, len(surface[i]))

		for j := range surface[i] {
			shallowCopyOfSurface[i][j] = &BreachHole{Address: surface[i][j].Address, IsFree: true}
		}

	}

	// Generate sequences based on surface
	var sequences = make([][]string, 0)
	var isRow = true
	var positionX = 0
	var positionY = 0
	var sequenceSize = length / count
	for i := 0; i < count; i++ {
		var resultingSequence = GenerateBreachSingleSequenceFromSurface(sequenceSize, shallowCopyOfSurface, positionX, positionY, isRow)

		sequences = append(sequences, resultingSequence.sequence)
		positionX = resultingSequence.positionX
		positionY = resultingSequence.positionY
		isRow = resultingSequence.isRow
	}

	// Randomize overlap of sequences
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < count-1; i++ {
		action := rand.Intn(3)
		switch action {
		case 0:
			// Keep it as it was
		case 1:
			// Add the first entry of the next sequence to the end of the current sequence
			var nextSequence = sequences[i+1]
			if len(nextSequence) > 0 {
				sequences[i] = append(sequences[i], nextSequence[0])
			}
		case 2:
			// Remove the last entry of the current sequence
			if len(sequences[i]) > 0 {
				sequences[i] = sequences[i][:len(sequences[i])-1]
			}
		}
	}
	rand.Shuffle(len(sequences), func(i, j int) { sequences[i], sequences[j] = sequences[j], sequences[i] })

	return sequences
}

func GenerateBreachBuffer(size int) []string {
	var breachBuffer = make([]string, size)
	for i := 0; i < size; i++ {
		breachBuffer[i] = BreachBufferFreeSlotSymbol
	}
	return breachBuffer
}
