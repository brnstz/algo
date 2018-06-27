package puzzles

func isMagic(sq [][]int, n int) bool {
	var i, j, k int

	// We have 2 diagonals plus n rows and cols
	sums := make([]int, 2+n*2)

	// Check first diagonal
	for i = 0; i < n; i++ {
		sums[k] += sq[i][i]
	}
	k++

	// Check second diagonal
	for i = 0; i < n; i++ {
		sums[k] += sq[i][n-1-i]
	}
	k++

	// Check horizontals
	for i = 0; i < n; i++ {
		for j = 0; j < n; j++ {
			sums[k] += sq[i][j]
		}
		k++
	}

	// Check verticals
	for j = 0; j < n; j++ {
		for i = 0; i < n; i++ {
			sums[k] += sq[i][j]
		}
		k++
	}

	for i := 1; i < len(sums); i++ {
		if sums[i] != sums[i-1] {
			return false
		}
	}

	return true
}

// MagicSquare returns the minimum cost required to convert s into a magic
// square.
func MagicSquare(sq [][]int) int {

	// Assume input is a square
	n := len(sq)

	// max := n * 9

	isMagic(sq, n)

	return 0
}
