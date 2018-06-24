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
