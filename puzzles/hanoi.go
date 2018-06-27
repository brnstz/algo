package puzzles

const (
	numHanoiPegs = 3
)

// HanoiPeg is a single peg on the Towers of Hanoi board. It holds zero
// or more rings, which are integers that represent their size. The left
// side of slice is the bottom, the right is the top.
type HanoiPeg struct {
	Rings []int
}

func moveRings(rings int, hp []HanoiPeg, source, target, aux int) [][]HanoiPeg {

	var moves [][]HanoiPeg
	move := make([]HanoiPeg, numHanoiPegs)

	// If there are no rings left, we are done
	if rings < 1 {
		return nil
	}

	// Recursively move rings from source to aux
	moves = append(moves, moveRings(rings-1, hp, source, aux, target)...)

	// Move a ring from source to the target
	hp[target].Rings = append(
		hp[target].Rings, hp[source].Rings[len(hp[source].Rings)-1],
	)
	hp[source].Rings = hp[source].Rings[:len(hp[source].Rings)-1]

	// Copy current state as a "move"
	for i := range hp {
		move[i].Rings = make([]int, len(hp[i].Rings))
		copy(move[i].Rings, hp[i].Rings)
	}

	moves = append(moves, move)

	// Recursively move ring from aux to target
	moves = append(moves, moveRings(rings-1, hp, aux, target, source)...)

	return moves
}

// TowersOfHanoi solves the Towers of Hanoi problem and returns a slice of
// slices representing the order state of the pegs.
func TowersOfHanoi(rings int) [][]HanoiPeg {
	hp := make([]HanoiPeg, numHanoiPegs)

	for i := 0; i < rings; i++ {
		hp[0].Rings = append(hp[0].Rings, i)
	}

	moves := moveRings(rings, hp, 0, 1, 2)

	return moves
}
