package algo

import (
	"testing"
)

func TestHanoi(t *testing.T) {
	moves := TowersOfHanoi(10)

	for _, move := range moves {
		// Ensure that there are no disc placed before one with
		// higher size
		for _, peg := range move {
			for i := 1; i < len(peg.Rings); i++ {
				if peg.Rings[i] < peg.Rings[i-1] {
					t.Fatalf("unexpected disc sequence: %v %v",
						peg.Rings[i-1], peg.Rings[i],
					)
				}
			}
		}
	}

	// The last move should have the 2nd peg in rings 0-9

	for i, ring := range moves[len(moves)-1][1].Rings {
		if i != ring {
			t.Fatalf("incorrect final move: %v", moves[len(moves)-1])
		}
	}
}
