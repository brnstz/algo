package algo

const (
	// MaxIntVal is the maximum possible int for this arch
	MaxIntVal = int(^uint(0) >> 1)

	// MinIntVal is the minimum possible int for this arch
	MinIntVal = -MaxIntVal - 1
)

// MaxInt returns maximum int in p
func MaxInt(p ...int) int {
	var max = MinIntVal

	for _, x := range p {
		if x > max {
			max = x
		}
	}

	return max
}

// MinInt returns minimum int in p
func MinInt(p ...int) int {
	var min = MaxIntVal

	for _, x := range p {
		if x < min {
			min = x
		}
	}

	return min
}
