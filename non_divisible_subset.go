package algo

// https://www.hackerrank.com/challenges/non-divisible-subset/problem

func isNonDivisibleSubset(k int, s []int) bool {

	for i := 0; i < len(s); i++ {
		for j := 0; j < len(s); j++ {

			if j == i {
				continue
			}

			if (s[i]+s[j])%k == 0 {
				return false
			}
		}
	}

	return true
}

// MaxNonDivisibleSubset returns the size of the maximal subset of s such that
// the sum of any two items is not divisible by k
func MaxNonDivisibleSubset(k int, s []int) int {
	if len(s) < 2 {
		return 0
	}

	if isNonDivisibleSubset(k, s) {
		return len(s)
	}

	// Otherwise, compute all possible combinations and find the best.
	maxSubset := 0

	for i := 0; i < len(s); i++ {
		// Remove exactly one element per iteration and find the
		// max subset
		newS := []int{}
		newS = append(newS, s[:i]...)
		newS = append(newS, s[i+1:]...)

		thisSubset := MaxNonDivisibleSubset(k, newS)

		if thisSubset > maxSubset {
			maxSubset = thisSubset
		}
	}

	return maxSubset
}
