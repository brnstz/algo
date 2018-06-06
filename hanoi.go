package algo

import "log"

const (
	numHanoiPegs = 3
)

// HanoiPeg FIXME
type HanoiPeg struct {
	Rings []int
}

func moveRings(disks int, hp []HanoiPeg, source, target, aux int) [][]HanoiPeg {

	var moves [][]HanoiPeg
	move := make([]HanoiPeg, numHanoiPegs)

	// If there are no disks left, we are done
	if disks < 1 {
		return nil
	}

	// Recursively move disks from source to aux
	moves = append(moves, moveRings(disks-1, hp, source, aux, target)...)

	// Move a disk from source to the target
	hp[target].Rings = append(hp[target].Rings, hp[source].Rings[len(hp[source].Rings)-1])
	hp[source].Rings = hp[source].Rings[:len(hp[source].Rings)-1]

	// FIXME: why doesn't this work? because move[i] isn't large enough
	for i := range hp {
		log.Printf("copying %v to %v", hp[i].Rings, move[i].Rings)
		copy(move[i].Rings, hp[i].Rings)
	}

	moves = append(moves, move)
	log.Printf("moves: %v", moves)

	// Recursively move disk from aux to target
	moves = append(moves, moveRings(disks-1, hp, aux, target, source)...)

	return moves
}

// TowersOfHanoi FIXME
func TowersOfHanoi(rings int) [][]HanoiPeg {
	hp := make([]HanoiPeg, numHanoiPegs)

	for i := 0; i < rings; i++ {
		hp[0].Rings = append(hp[0].Rings, i)
	}

	moves := moveRings(rings, hp, 0, 1, 2)

	return moves
}
