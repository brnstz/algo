package algo

func gsHelper(subpattern string, grid [][]byte, x, y int) bool {

	if len(subpattern) < 1 {
		return true
	}

	// If we aren't starting with this byte, then not possible
	if grid[x][y] != subpattern[0] {
		return false
	}

	subpattern = subpattern[1:]

	if x+1 < len(grid) {
		if gsHelper(subpattern, grid, x+1, y) {
			return true
		}
	}

	if y+1 < len(grid[x]) {
		if gsHelper(subpattern, grid, x, y+1) {
			return true
		}
	}

	if x-1 > 0 {
		if gsHelper(subpattern, grid, x-1, y) {
			return true
		}
	}

	if y-1 > 0 {
		if gsHelper(subpattern, grid, x, y-1) {
			return true
		}
	}

	return false
}

// GridSearch determines if pattern can be found in the grid of bytes
// by moving one space directly left, right, up, or down. Diagonal moves
// are not allowed. Repeat moves *are* allowed.
func GridSearch(pattern string, grid [][]byte) bool {

	if len(grid) < 1 {
		return false
	}

	// Assume that all rows are the same length.
	if len(grid[0]) < 1 {
		return false
	}

	for x := 0; x < len(grid); x++ {
		for y := 0; y < len(grid[x]); y++ {
			if gsHelper(pattern, grid, x, y) {
				return true
			}
		}
	}

	return false
}
