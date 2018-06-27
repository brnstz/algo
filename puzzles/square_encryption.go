package puzzles

import (
	"math"
)

// SquareEncryption returns a space delimited string by:
//    - Removing all spaces from the string
//    - Encoding the text as grid of LrxLc
//    - Where L is between the floor/ceiling of the square root
//        of the size of the string with the spaces removed
//        and Lr <= Lc
//    - Words are written columnwise
//    - Rows are returned as a space-delimited string
func SquareEncryption(s string) string {
	var (
		char                    rune
		in, out                 string
		i, j, index, rows, cols int
		root                    float64
	)

	for _, char = range s {
		switch char {
		case ' ':
		default:
			in += string(char)
		}
	}

	root = math.Sqrt(float64(len(s)))

	rows = int(math.Floor(root))
	cols = int(math.Ceil(root))

	for i = 0; i < rows; i++ {
		for j = 0; j < cols; j++ {
			index = rows*j + i
			if index < len(in) {
				out += string(in[index])
			}
		}
		if i != rows-1 {
			out += " "
		}
	}

	return out
}

/*

	hellothere

	3 rows
	4 colums

    012
0   hel
1	lot
2	her
3   e

[0,0], [0,1], [0, 2], [0,3]
[1,0], [1,1], [1, 2], [1,3]
[2,0], [2,1]

0, 3, 6, 9

hellothere
0123456789
*/
