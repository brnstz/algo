package algo

import "fmt"

// https://www.hackerrank.com/challenges/organizing-containers-of-balls/problem

func swappable(container [][]int, n int, assigned int, totals []int, ballAssignment []int, containerUsed map[int]bool) bool {
	// totals[i] = the total number of balls of type i
	// bA[i] = ball i is assigned to container bA[i]
	// cA[i] = container i is assigned to ball cA[i]

	fmt.Printf("%v %v %v %v %v %v\n\n",
		container, n, assigned, totals, ballAssignment, containerUsed,
	)

	var (
		deficit, i, j int
	)

	if assigned == n {
		for i = 0; i < n; i++ {
			deficit += totals[i] - container[ballAssignment[i]][i]
		}

		fmt.Printf("what is the deficit? %v\n", deficit)

		return deficit%2 == 0
	}

	// Assign the next ball
	assigned++

	// Assign the ball to all possible containers
	for j = 0; j < n; j++ {
		if containerUsed[j] {
			continue
		}

		newCU := map[int]bool{}
		for k, v := range containerUsed {
			newCU[k] = v
		}
		newCU[j] = true

		newBA := make([]int, len(ballAssignment)+1)
		copy(newBA, ballAssignment)
		fmt.Printf("%v %v %v\n", newBA, ballAssignment, assigned)
		newBA[assigned-1] = j

		if swappable(container, n, assigned, totals, newBA, newCU) {
			return true
		}

	}

	/*
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
	*/

	return false
}

// ContainerSwap FIXME
func ContainerSwap(container [][]int) bool {
	var (
		i, j, n int

		totals []int
	)

	// Assume square matrix
	n = len(container)

	totals = make([]int, n)

	for i = 0; i < n; i++ {
		for j = 0; j < n; j++ {
			totals[j] += container[i][j]
		}
	}
	containerUsed := map[int]bool{}

	return swappable(container, n, 0, totals, nil, containerUsed)
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
