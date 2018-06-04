package algo

const bitLen = 64

// BruteForceBitParity computes the parity of a uint64 using an O(n) algorithm,
// returning true if the parity is even, false if odd.
func BruteForceBitParity(x uint64) bool {
	var (
		result int
		i      uint8
	)

	for i = 0; i < bitLen; i++ {
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

	for i = bitLen / 2; i > 0; i = i / 2 {
		x ^= x >> i
	}

	return x&0x1 == 0
}

// BitSwap swaps the bits i and j in x
func BitSwap(x int64, i, j uint) int64 {
	// If bits are the same, nothing to do.
	if ((x >> i) & 1) == ((x >> j) & 1) {
		return x
	}

	// Otherwise, we can just independently flip the bits
	x ^= (1 << i) | (1 << j)

	return x
}

// BitReverse returns x with its bits reversed
func BitReverse(x uint64) uint64 {
	var i, j uint

	// Swap each bit by doing two at a time, only go up to
	// bitLen / 2
	for i = 0; i < bitLen/2; i++ {

		// j is the high bit
		j = bitLen - 1 - i

		// Does the bit need to be swapped? If not, continue.
		if ((x >> i) & 1) == ((x >> j) & 1) {
			continue
		}

		// Otherwise, we just need to flip both.
		x ^= (1 << i) | (1 << j)
	}

	return x
}
