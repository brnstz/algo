package puzzles

import (
	"math"
)

/*
SquareEncryption returns a space delimited string by:
    - Removing all spaces from the string
    - Encoding the text as grid of LrxLc
    - Where L is between the floor/ceiling of the square root
        of the size of the string with the spaces removed
        and Lr <= Lc
    - Words are written columnwise
    - Rows are returned as a space-delimited string

	s=hellothere

	3 rows
	4 columns

		0123
	0   hell
	1   othe
	2   re

	returns hor ete lh le

	print these coords in order
	[0, 0], [1, 0], [2, 0]
	[1, 0], [1, 1], [1, 2]
	[2, 0], [2, 1]
	[3, 0], [3, 1]

	which maps to these indexes in space-removed string
	0, 4, 8
	1, 5, 9
	2, 6
	3, 8

	hellothere
	0123456789
*/
func SquareEncryption(s string) string {
	var (
		char             rune
		in, out          string
		i, j, rows, cols int
		root             float64
	)

	// Create in string with no spaces
	for _, char = range s {
		switch char {
		case ' ':
		default:
			in += string(char)
		}
	}

	// Get the float square root
	root = math.Sqrt(float64(len(s)))

	// Get the floor and ceiling, ensuring that
	// rows is <= cols
	rows = int(math.Floor(root))
	cols = int(math.Ceil(root))

	// If this doesn't give us enough space, increase rows
	if rows*cols < len(s) {
		rows++
	}

	// Go through every column
	for i = 0; i < cols; i++ {
		// Print the corresponding string index for each row of the column
		for j = i; j < len(in); j += cols {
			out += string(in[j])
		}

		// If it's not the last one, append a space
		if i != cols-1 {
			out += " "
		}
	}

	return out
}
