package breachModel

import (
	"math/rand"
	"time"
)

// The addresses of the breach holes that are known to hackers
var knownBreachHoleAddresses = [...]string{"1C", "55", "7A", "BD", "E9", "FF"}

// Represents a breach hole with a known address that can be used to hack into a system
type BreachHole struct {
	Address string
	IsFree bool
}

// Generates an array of random breach hole addresses with a length [2, maxLength]
func GenerateBreachSequence(maxLength int) []string {
	rand.Seed(time.Now().UnixNano())
	var size = 0
	for size < 2 {
		size = rand.Intn(maxLength + 1) 
	}
	var breachSequence = make([]string, size)
	for i := 0; i < size; i++ {
		randItemIndex := rand.Intn(len(knownBreachHoleAddresses))
		breachSequence[i] = knownBreachHoleAddresses[randItemIndex]
	}
	return breachSequence
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