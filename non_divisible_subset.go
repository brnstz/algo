package algo

import (
	"strconv"
	"strings"
)

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

// mndsKey returns a key for a cached version of a subset solution, assuming
// that repeated attempts for s are always in the same order
func mndsKey(s []int) string {
	k := make([]string, len(s))
	for i, v := range s {
		k[i] = strconv.Itoa(v)
	}

	return strings.Join(k, "|")
}

func mnds(k int, s []int, cache map[string]int) int {
	var (
		key    string
		value  int
		exists bool
	)

	// Check the cache first
	key = mndsKey(s)
	value, exists = cache[key]
	if exists {
		return value
	}

	if len(s) < 2 {
		cache[key] = 1
		return 1
	}

	if isNonDivisibleSubset(k, s) {
		cache[key] = len(s)
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

		thisSubset := mnds(k, newS, cache)

		if thisSubset > maxSubset {
			maxSubset = thisSubset
		}
	}

	cache[key] = maxSubset
	return maxSubset
}

// MaxNonDivisibleSubset returns the size of the maximal subset of s such that
// the sum of any two items is not divisible by k
func MaxNonDivisibleSubset(k int, s []int) int {
	cache := map[string]int{}
	return mnds(k, s, cache)
}

// MaxNonDivisibleSubsetIterative returns the size of the maximal subset of s
// such that the sum of any two items is not divisible by k by pairing lists of
// remainders iteratively.
func MaxNonDivisibleSubsetIterative(k int, s []int) int {
	var remainder int

	// Remainders is a mapping of the remainder of dividing a value in s by k,
	// mapped to list of values in s that have that remainder
	remainders := make([][]int, k+1)

	// Create the mapping as described above
	for _, v := range s {
		remainder = v % k
		remainders[remainder] = append(remainders[remainder], v)
	}

	// Each remainder has a corresponding entry in the list: the value
	// that it needs to add up to k. Iterate through all of them and
	// choose the bigger corresponding entry each time.
	optimalSubset := []int{}
	for i := 1; i <= k/2; i++ {

		// If k is even, we want to skip the entry that is k/2.  Adding > 1
		// value from this will result in a divisible pair. Instead, just add 1 from
		// this list if it exists.
		if i == k-i {
			if len(remainders[i]) > 0 {
				optimalSubset = append(optimalSubset, remainders[i][0])
			}
			continue
		}

		// Choose which of the remainder pair to include in the optimal subset
		if len(remainders[i]) > len(remainders[k-i]) {
			optimalSubset = append(optimalSubset, remainders[i]...)
		} else {
			optimalSubset = append(optimalSubset, remainders[k-i]...)
		}
	}

	// We can also include exactly one entry that is evenly divisble
	if len(remainders[0]) > 0 {
		optimalSubset = append(optimalSubset, remainders[0][0])
	}

	return len(optimalSubset)
}
