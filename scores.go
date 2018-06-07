package algo

import "fmt"

var (
	// AllPossiblePoints FIXME
	AllPossiblePoints = []int{7, 6, 3, 2}
	pointCache        = map[int]int{}
)

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
