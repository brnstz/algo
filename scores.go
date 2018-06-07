package algo

import "fmt"

// PointCombinationSolver FIXME
type PointCombinationSolver struct {
	possiblePoints []int
	comboCache     map[int]int
}

// NewPointCombinationSolver FIXME
func NewPointCombinationSolver(possiblePoints []int, points int) *PointCombinationSolver {
	pcs := &PointCombinationSolver{
		possiblePoints: possiblePoints,
		points:         points,
		comboCache:     map[int]int{},
	}
	return pcs
}

func (pcs *PointCombinationSolver) _solve(points int) int {
	var (
		combos int
		exists bool
	)

	combos, exists = comboCache[points]
	if exists {
		return combos
	}

}

// Solve FIXME
func (pcs *PointCombinationSolver) Solve() {

}

// PointCombinations returns the number of ways we can arrive an points
// given possiblePoints
func PointCombinations(points int, possiblePoints []int) int {
	fmt.Printf("points: %v\n", points)
	var (
		combinations    int
		combinationsNow int
		exists          bool
		pointsNow       int
	)

	// Check cache first
	combinations, exists = pointCache[points]
	if exists {
		return combinations
	}

	for _, poss := range possiblePoints {
		pointsNow = points
		for pointsNow-poss >= 0 {
			pointsNow = pointsNow - poss

			if pointsNow == 0 {
				fmt.Printf("now: %v %v ++\n", pointsNow, poss)
				combinations++
				continue
			}

			if len(possiblePoints) > 1 {
				combinationsNow = PointCombinations(
					pointsNow,
					possiblePoints[:len(possiblePoints)-1],
				)
				combinations += combinationsNow
				fmt.Printf("combo: %v %v %v %v ++\n",
					pointsNow, poss, combinations, combinationsNow)
			}
		}
	}

	pointCache[points] = combinations

	fmt.Printf("total: %v => %v", points, combinations)
	fmt.Println()
	fmt.Println()
	return combinations
}
