package algo

import "log"

// https://www.hackerrank.com/challenges/organizing-containers-of-balls/problem

func swappable(container [][]int, n int, totals []int, bA []int, cA []int) bool {
	// totals[i] = the total number of balls of type i
	// assignments[i] = the container where balls of type i are stored

	var (
		deficit, i, j int
	)

	if bA[n-1] != -1 {
		for i = 0; i < n; i++ {
			deficit = totals[i] - container[bA[i]][i]
		}

		return deficit%2 == 0
	}

	for i = 0; i < n; i++ {
		for j = 0; j < n; j++ {
			if bA[i] == -1 && cA[j] == -1 {
				newBA := make([]int, n)
				newCA := make([]int, n)
				copy(newBA, bA)
				copy(newCA, cA)

				newBA[i] = j
				newBA[j] = i

				if swappable(container, n, totals, newBA, newCA) {
					return true
				}
			}
		}
	}

	return false
}

// ContainerSwap FIXME
func ContainerSwap(container [][]int) bool {
	var (
		i, j, n int

		totals, bA, cA []int
	)

	// Assume square matrix
	n = len(container)

	bA = make([]int, n)
	cA = make([]int, n)
	totals = make([]int, n)

	for i = 0; i < n; i++ {
		log.Printf("%v %v", len(bA), i)
		bA[i] = -1
		cA[i] = -1
		for j = 0; j < n; j++ {
			totals[j] += container[i][j]
		}
	}

	return swappable(container, n, totals, bA, cA)
}

// - each container contains only balls of the same type
// - no two balls of the same type are located in different
//   containers

// for each row:
//   - one column is >= 0
//   - all other columns are == 0
//   - no two rows have same column >= 0

// brute force solution
//	- find all possible combinations of rows and columns
//		(e.g., assign a column to be >= to a row)
//  - recursively perform all moves leading to that row
//		matching the conditions

// more optimized solution
// - don't actually solve the problem, just figure out if it's
//	 possible
//
// - what are the cases where it wouldn't be possible?
