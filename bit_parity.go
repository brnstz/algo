package algo

const bitParityLen = 64

// BruteForceBitParity computes the parity of a uint64
func BruteForceBitParity(x uint64) int {
	var (
		result int
		i      uint8
	)

	for i = 0; i < bitParityLen; i++ {
		result += int(x & 1)
		x = x >> 1
	}

	return result
}
