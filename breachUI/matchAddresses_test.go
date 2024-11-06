package breachUI

import (
	"main/breachModel"
	"testing"
)

func TestMatchAddresses(t *testing.T) {
	tests := []struct {
		name            string
		sequence        []string
		breachBuffer    []string
		expectedOffset  int
		expectedMatches int
	}{
		{
			name:            "Full match at start",
			sequence:        []string{"A", "B", "C"},
			breachBuffer:    []string{"A", "B", "C", breachModel.BreachBufferFreeSlotSymbol, breachModel.BreachBufferFreeSlotSymbol},
			expectedOffset:  0,
			expectedMatches: 3,
		},
		{
			name:            "Full match with repeating start characters",
			sequence:        []string{"A", "B"},
			breachBuffer:    []string{"A", "A", "A", "A", "B", "A"},
			expectedOffset:  3,
			expectedMatches: 2,
		},
		{
			name:            "Full match with repeating start sequence characters",
			sequence:        []string{"A", "B", "C"},
			breachBuffer:    []string{"A", "B", "A", "B", "C", "A"},
			expectedOffset:  2,
			expectedMatches: 3,
		},
		{
			name:            "Full match with repeating characters and same characters",
			sequence:        []string{"A", "A"},
			breachBuffer:    []string{"A", "B", "A", "A", "A", "A"},
			expectedOffset:  2,
			expectedMatches: 2,
		},
		{
			name:            "Full match in the middle",
			sequence:        []string{"B", "C"},
			breachBuffer:    []string{"A", "A", "B", "C", "B", "C"},
			expectedOffset:  2,
			expectedMatches: 2,
		},
		{
			name:            "Full match at the end",
			sequence:        []string{"C", "D"},
			breachBuffer:    []string{"A", "B", "C", "D", breachModel.BreachBufferFreeSlotSymbol, breachModel.BreachBufferFreeSlotSymbol},
			expectedOffset:  2,
			expectedMatches: 2,
		},
		{
			name:            "No match",
			sequence:        []string{"X", "Y", "Z"},
			breachBuffer:    []string{"A", "B", "C", "A", "A", "A"},
			expectedOffset:  0,
			expectedMatches: 0,
		},
		{
			name:            "Partial match with possibility for complete match",
			sequence:        []string{"A", "B", "C"},
			breachBuffer:    []string{"A", "A", "A", "A", "B", breachModel.BreachBufferFreeSlotSymbol},
			expectedOffset:  3,
			expectedMatches: 2,
		},
		{
			name:            "Partial match with no possibility for complete match",
			sequence:        []string{"A", "B", "C", "D"},
			breachBuffer:    []string{"A", "A", "A", "A", "B", breachModel.BreachBufferFreeSlotSymbol},
			expectedOffset:  3,
			expectedMatches: 2,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			offset, matchesCount := matchAddresses(tt.sequence, tt.breachBuffer)
			if offset != tt.expectedOffset || matchesCount != tt.expectedMatches {
				t.Errorf("matchAddresses() = (%d, %d), want (%d, %d)", offset, matchesCount, tt.expectedOffset, tt.expectedMatches)
			}
		})
	}
}
