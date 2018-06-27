package algo

// https://www.hackerrank.com/challenges/organizing-containers-of-balls/problem

// ContainerSwap determines if, given container[i][j] all balls of type j can
// be put into a single container i by making a series of swaps of balls
// between containers.
func ContainerSwap(container [][]int) bool {
	var (
		i, j, n int

		containerTotals []int
		ballTotals      []int

		counts map[int]bool
	)

	// Assume square matrix
	n = len(container)

	// containerTotals[i] is the number of all balls in container i
	containerTotals = make([]int, n)

	// ballTotals[j] is the number of ball type j in all container
	ballTotals = make([]int, n)

	// counts is a mapping of total values that should exist in both
	// containerTotals and ballTotals
	counts = map[int]bool{}

	for i = 0; i < n; i++ {
		for j = 0; j < n; j++ {
			containerTotals[i] += container[i][j]
			ballTotals[j] += container[i][j]
		}
	}

	for i = 0; i < n; i++ {
		counts[ballTotals[i]] = true
	}

	for i = 0; i < n; i++ {
		if !counts[containerTotals[i]] {
			return false
		}
	}

	return true
}
