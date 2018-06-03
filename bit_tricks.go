package algo

const bitParityLen = 64

// BruteForceBitParity computes the parity of a uint64 using an O(n) algorithm,
// returning true if the parity is even, false if odd.
func BruteForceBitParity(x uint64) bool {
	var (
		result int
		i      uint8
	)

	for i = 0; i < bitParityLen; i++ {
		result += int(x & 1)
		x = x >> 1
	}

	return result%2 == 0
}

// BitParity computes the parity of a uint64 using an O(log n) algorithm,
// return true if the parity is even, false if odd.
func BitParity(x uint64) bool {
	var (
		i uint8
	)

	for i = bitParityLen / 2; i > 0; i = i / 2 {
		x ^= x >> i
	}

	return x&0x1 == 0
}

// BitSwap swaps the bits i and j in x
func BitSwap(x int, i, j uint) int {
	// If bits are the same, nothing to do.
	if ((x >> i) & 1) == ((x >> j) & 1) {
		return x
	}

	// Otherwise, we can just independently swap the bits
	x ^= (1 << i) | (1 << j)

	return x
}
